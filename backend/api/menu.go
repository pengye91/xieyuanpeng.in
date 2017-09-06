// Todo: Add MenuItems API. This should the root of all front-end classification.
// This should also be reflected in front-end router and vuex.
// four entities in redis:
// sideMenuItem: {
// 'blog': {'python': 'python', 'golang': 'golang', 'django': 'django', 'miscellaneous': '杂'},
// 'photography': {'project-1': '项目1', 'project-2': '项目2'},
// 'contact':{'github': 'github', 'wechat': 'wechat'}
// }
//
// menuItems: {
//    'blog': {
//      ref: 'blog',
//      name: '博客',
//      sideMenuItems: {'python': 'python', 'golang': 'golang', 'django': 'django', 'miscellaneous': '杂'},
//      adminSideMenuItems: adminSideMenuItem
//    },
//    'photography': {
//      ref: 'photography',
//      name: '摄影',
//      sideMenuItems: {'project-1': '项目1', 'project-2': '项目2'},
//      adminSideMenuItems: adminSideMenuItem
//    },
//    'contact': {
//      ref: 'contact',
//      name: '联系我',
//      sideMenuItems: {'github': 'github', 'wechat': 'wechat'},
//      adminSideMenuItems: adminSideMenuItem
//    }
//  }
// }
//
// menuID: "XXXXXXXXXXXXXXXXXXXXXXXX"
//
// adminSideMenuItems: {
//	"blog": {
//		'all': '所有',
//        	'upload': '上传',
//        	'with-selected': '选中操作'
//	},
//	"photography": {
//		'all': '所有',
//        	'upload': '上传',
//        	'with-selected': '选中操作'
//	},
//	'contact': {
//		'all': '所有',
//        	'upload': '上传',
//        	'with-selected': '选中操作'
//	}
// }
//
//
// But only menuItems in mongo.

package api

import (
	"encoding/json"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-errors"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/cache"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type MenuApi struct {
	*gin.Context
}

func (this MenuApi) GetMenu(ctx *gin.Context) {
	getMenuItems(ctx, "menuItems")
}

func (this MenuApi) GetSideMenu(ctx *gin.Context) {
	getMenuItems(ctx, "sideMenuItems")
}

func (this MenuApi) GetAdminSideMenu(ctx *gin.Context) {
	getMenuItems(ctx, "adminSideMenuItems")
}

func getMenuItems(ctx *gin.Context, itemType string) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	redisConn := cache.GlobalUserRedisPool.Get()
	defer redisConn.Close()

	var (
		menu               models.Menu
		sideMenuItems      = make(models.SideMenuItems)
		adminSideMenuItems = make(models.AdminSideMenuItems)
		menuItems          = make(map[string]models.MenuItem)
	)

	reply, getMenuItemsError := redis.Values(redisConn.Do("HMGET", "menu",
		"menuID",
		itemType,
	))
	if getMenuItemsError != nil {
		log.LoggerSugar.Errorw("menu getMentItems HMGET menu Error",
			"module", "redis",
			"error", getMenuItemsError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	if reply[1] != nil {
		// Always Unmarshal to a pointer.
		switch itemType {
		case "menuItems":
			unmarshalError := json.Unmarshal(reply[1].([]byte), &menuItems)
			if unmarshalError != nil {
				log.LoggerSugar.Errorw("menu getMentItems json.Unmarshal Error",
					"module", "application",
					"error", unmarshalError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			log.LoggerSugar.Infow("menu getMentItems Get menuItems from Redis succeed",
				"module", "redis",
			)

			ctx.JSON(http.StatusOK, menuItems)
			return
		case "sideMenuItems":
			unmarshalError := json.Unmarshal(reply[1].([]byte), &sideMenuItems)
			if unmarshalError != nil {
				log.LoggerSugar.Errorw("menu getMentItems json.Unmarshal Error",
					"module", "application",
					"error", unmarshalError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			log.LoggerSugar.Infow("menu getMentItems Get sideMenuItems from Redis succeed",
				"module", "redis",
			)

			ctx.JSON(http.StatusOK, sideMenuItems)
			return
		case "adminSideMenuItems":
			unmarshalError := json.Unmarshal(reply[1].([]byte), &adminSideMenuItems)
			if unmarshalError != nil {
				log.LoggerSugar.Errorw("menu getMentItems json.Unmarshal Error",
					"module", "application",
					"error", unmarshalError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			log.LoggerSugar.Infow("menu getMentItems Get adminSideMenuItems from Redis succeed",
				"module", "redis",
			)

			ctx.JSON(http.StatusOK, adminSideMenuItems)
			return
		}
	} else {
		log.LoggerSugar.Infow("menu getMentItems Get items from Redis failed, trying to get from db",
			"module", "redis",
		)
		if getMenuErr := Db.C("menu").Find(nil).One(&menu); getMenuErr != nil {
			log.LoggerSugar.Errorw("menu getMentItems Error",
				"module", "mongo",
				"error", getMenuErr,
			)
			ctx.JSON(http.StatusInternalServerError, models.Err("5"))
			return
		}
		switch itemType {
		case "menuItems":
			dbMenuItems, jsonMarshalError := json.Marshal(menu.MenuItems)
			if jsonMarshalError != nil {
				log.LoggerSugar.Errorw("menu getMentItems Error: jsonMarshalError",
					"module", "application json marshal",
					"error", jsonMarshalError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			log.LoggerSugar.Infow("menu getMentItems trying to get menu from db succeed",
				"module", "mongo",
			)

			_, hmsetError := redis.String(redisConn.Do("HMSET", "menu",
				"menuItems", dbMenuItems,
			))
			if hmsetError != nil {
				log.LoggerSugar.Errorw("menu getMentItems HMSET menu Error",
					"module", "redis",
					"error", hmsetError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			ctx.JSON(http.StatusOK, menu.MenuItems)
			return
		case "sideMenuItems":
			for menuItemName, menuItem := range menu.MenuItems {
				sideMenuItems[menuItemName] = menuItem.SideMenuItems
			}
			dbSideMenuItems, jsonMarshalError := json.Marshal(sideMenuItems)
			if jsonMarshalError != nil {
				log.LoggerSugar.Errorw("menu getMentItems Error: jsonMarshalError",
					"module", "application json marshal",
					"error", jsonMarshalError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			log.LoggerSugar.Infow("menu getMenutItems trying to get menu from db succeed",
				"module", "mongo",
			)

			_, hmsetError := redis.String(redisConn.Do("HMSET", "menu",
				"sideMenuItems", dbSideMenuItems,
			))
			if hmsetError != nil {
				log.LoggerSugar.Errorw("menu getMenutItems HMSET menu Error",
					"module", "redis",
					"error", hmsetError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			ctx.JSON(http.StatusOK, sideMenuItems)
			return
		case "adminSideMenuItems":
			for menuItemName, menuItem := range menu.MenuItems {
				adminSideMenuItems[menuItemName] = menuItem.AdminSideMenuItems
			}
			dbAdminSideMenuItems, jsonMarshalError := json.Marshal(adminSideMenuItems)
			if jsonMarshalError != nil {
				log.LoggerSugar.Errorw("menu getMentItems Error: jsonMarshalError",
					"module", "application json marshal",
					"error", jsonMarshalError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			log.LoggerSugar.Infow("menu getMenutItems trying to get menu from db succeed",
				"module", "mongo",
			)

			_, hmsetError := redis.String(redisConn.Do("HMSET", "menu",
				"adminSideMenuItems", dbAdminSideMenuItems,
			))
			if hmsetError != nil {
				log.LoggerSugar.Errorw("menu getMenutItems HMSET menu Error",
					"module", "redis",
					"error", hmsetError,
				)
				ctx.JSON(http.StatusInternalServerError, models.Err("5"))
				return
			}
			ctx.JSON(http.StatusOK, adminSideMenuItems)
			return
		}
	}
}

func (this MenuApi) PutSideMenuItem(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	redisConn := cache.GlobalUserRedisPool.Get()
	defer redisConn.Close()

	var (
		sideMenuItems   = make(models.SideMenuItems)
		cachedMenuItems = make(map[string]models.MenuItem)
	)

	bindJsonErr := ctx.BindJSON(&sideMenuItems)
	if bindJsonErr != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem bindJson Error",
			"module", "application",
			"error", bindJsonErr,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("2"))
		return
	}

	menuItemName := ctx.Query("menu-item")
	if menuItemName == "" {
		log.LoggerSugar.Warnw("menu PutAdminSideMenuItem ctx.Query Warn",
			"module", "application: ctx.Query",
			"warn", "no query menu-item get",
		)
		ctx.JSON(http.StatusBadRequest, errors.New("No menu-item in query params"))
		return
	}

	hMgetreply, getMenuError := redis.Values(redisConn.Do("HMGET", "menu",
		"menuID",
		"menuItems",
		"sideMenuItems",
	))
	if getMenuError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem HMGET menu Error",
			"module", "redis",
			"error", getMenuError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	jsonUnmarshalError := json.Unmarshal(hMgetreply[1].([]byte), &cachedMenuItems)
	if jsonUnmarshalError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem json.Unmarshal Error",
			"module", "application: json.Unmarshal",
			"error", jsonUnmarshalError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	// This is stupid but necessary.
	tmpMenuItem := cachedMenuItems[menuItemName]
	tmpMenuItem.SideMenuItems = sideMenuItems[menuItemName]
	cachedMenuItems[menuItemName] = tmpMenuItem

	updatedMenuItems, jsonMarshalError := json.Marshal(cachedMenuItems)
	updatedSideMenuItems, jsonMarshalError0 := json.Marshal(sideMenuItems)
	if (jsonMarshalError != nil) || (jsonMarshalError0 != nil) {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem json.Marshal Error",
			"module", "application: json.Marshal",
			"error", jsonMarshalError,
			"error0", jsonMarshalError0,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	_, hmsetError := redis.String(redisConn.Do("HMSET", "menu",
		"menuItems", updatedMenuItems,
		"sideMenuItems", updatedSideMenuItems,
	))
	if hmsetError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem HMSET menu Error",
			"module", "redis",
			"error", hmsetError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	updateSideMenuItem := bson.M{
		"$set": bson.M{
			"menuItems." + menuItemName + ".sideMenuItems": sideMenuItems,
		},
	}

	if str, ok := hMgetreply[0].([]byte); ok {
		fmt.Println(str)
	} else {
		fmt.Println("NOT OK")
	}
	updateSideMenuItemsError := Db.C("menu").UpdateId(bson.ObjectIdHex(string(hMgetreply[0].([]uint8))), updateSideMenuItem)
	if updateSideMenuItemsError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem mongo UpdateId Error",
			"module", "mongo",
			"error", updateSideMenuItemsError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	log.LoggerSugar.Infow("menu PutAdminSideMenuItem mongo UpdateMenu Success",
		"info", "update sideMenuItem in menu success.",
	)
	ctx.JSON(http.StatusOK, gin.H{"sideMenuItems": sideMenuItems, "menuItems": cachedMenuItems})

}

func (this MenuApi) PutAdminSideMenuItem(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	redisConn := cache.GlobalUserRedisPool.Get()
	defer redisConn.Close()

	var (
		adminSideMenuItems = make(models.AdminSideMenuItems)
		cachedMenuItems    = make(map[string]models.MenuItem)
	)

	bindJsonErr := ctx.BindJSON(&adminSideMenuItems)
	if bindJsonErr != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem bindJson Error",
			"module", "application",
			"error", bindJsonErr,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("2"))
		return
	}

	menuItemName := ctx.Query("menu-item")
	if menuItemName == "" {
		log.LoggerSugar.Warnw("menu PutAdminSideMenuItem ctx.Query Warn",
			"module", "application: ctx.Query",
			"warn", "no query menu-item get",
		)
		ctx.JSON(http.StatusBadRequest, errors.New("No menu-item in query params"))
		return
	}

	hMgetreply, getMenuError := redis.Values(redisConn.Do("HMGET", "menu",
		"menuID",
		"menuItems",
		"adminSideMenuItems",
	))
	if getMenuError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem HMGET menu Error",
			"module", "redis",
			"error", getMenuError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	jsonUnmarshalError := json.Unmarshal(hMgetreply[1].([]byte), &cachedMenuItems)
	if jsonUnmarshalError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem json.Unmarshal Error",
			"module", "application: json.Unmarshal",
			"error", jsonUnmarshalError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	// This is stupid but necessary.
	tmpMenuItem := cachedMenuItems[menuItemName]
	tmpMenuItem.AdminSideMenuItems = adminSideMenuItems[menuItemName]
	cachedMenuItems[menuItemName] = tmpMenuItem

	updatedMenuItems, jsonMarshalError := json.Marshal(cachedMenuItems)
	updatedAdminSideMenuItems, jsonMarshalError0 := json.Marshal(adminSideMenuItems)
	if (jsonMarshalError != nil) || (jsonMarshalError0 != nil) {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem json.Marshal Error",
			"module", "application: json.Marshal",
			"error", jsonMarshalError,
			"error0", jsonMarshalError0,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	_, hmsetError := redis.String(redisConn.Do("HMSET", "menu",
		"menuItems", updatedMenuItems,
		"adminSideMenuItems", updatedAdminSideMenuItems,
	))
	if hmsetError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem HMSET menu Error",
			"module", "redis",
			"error", hmsetError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	updateAdminSideMenuItem := bson.M{
		"$set": bson.M{
			"menuItems." + menuItemName + ".adminSideMenuItems": adminSideMenuItems,
		},
	}

	updateAdminSideMenuItemsError := Db.C("menu").UpdateId(bson.ObjectIdHex(hMgetreply[0].(string)), updateAdminSideMenuItem)
	if updateAdminSideMenuItemsError != nil {
		log.LoggerSugar.Errorw("menu PutAdminSideMenuItem mongo UpdateId Error",
			"module", "mongo",
			"error", updateAdminSideMenuItemsError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	log.LoggerSugar.Infow("menu PutAdminSideMenuItem mongo UpdateMenu Success",
		"info", "update sideMenuItem in menu success.",
	)
	ctx.JSON(http.StatusOK, gin.H{"adminSideMenuItems": adminSideMenuItems, "menuItems": cachedMenuItems})
}

// request body is just one menuItem, not the whole menu.
func (this MenuApi) PutMenuItem(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	redisConn := cache.GlobalUserRedisPool.Get()

	defer Db.Close()
	defer redisConn.Close()
	var (
		cachedSideMenuItems      = make(models.SideMenuItems)
		cachedAdminSideMenuItems = make(models.AdminSideMenuItems)
		menuItem                 = make(map[string]models.MenuItem)
		cachedMenuItems          = make(map[string]models.MenuItem)
	)
	if err := ctx.BindJSON(&menuItem); err != nil {
		log.LoggerSugar.Errorw("menu PutMenuItem bindJson Error",
			"module", "application: BindJson",
			"error", err,
		)
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
		return
	}

	menuItemName := ctx.Query("menu-item")
	if menuItemName == "" {
		log.LoggerSugar.Warnw("menu PutMenuItem ctx.Query Warn",
			"module", "application: ctx.Query",
			"warn", "no query menu-item get",
		)
		ctx.JSON(http.StatusBadRequest, errors.New("No menu-item in query params"))
		return
	}

	hMgetreply, getMenuError := redis.Values(redisConn.Do("HMGET", "menu",
		"menuID",
		"menuItems",
		"sideMenuItems",
		"adminSideMenuItems",
	))
	if getMenuError != nil {
		log.LoggerSugar.Errorw("menu PutMenuItem HGET menu Error",
			"module", "redis",
			"error", getMenuError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	jsonUnmarshalError := json.Unmarshal(hMgetreply[1].([]byte), &cachedMenuItems)
	jsonUnmarshalError0 := json.Unmarshal(hMgetreply[2].([]byte), &cachedSideMenuItems)
	jsonUnmarshalError1 := json.Unmarshal(hMgetreply[3].([]byte), &cachedAdminSideMenuItems)
	if (jsonUnmarshalError != nil) || (jsonUnmarshalError0 != nil) || (jsonUnmarshalError1 != nil) {
		log.LoggerSugar.Errorw("menu PutMenuItem json.Unmarshal Error",
			"module", "application: json.Unmarshal",
			"error", jsonUnmarshalError,
			"error0", jsonUnmarshalError0,
			"error1", jsonUnmarshalError1,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	cachedMenuItems[menuItemName] = menuItem[menuItemName]
	cachedSideMenuItems[menuItemName] = menuItem[menuItemName].SideMenuItems
	cachedAdminSideMenuItems[menuItemName] = menuItem[menuItemName].AdminSideMenuItems

	updatedMenuItems, jsonMarshalError := json.Marshal(cachedMenuItems)
	updatedSideMenuItems, jsonMarshalError0 := json.Marshal(cachedSideMenuItems)
	updatedAdminSideMenuItems, jsonMarshalError1 := json.Marshal(cachedAdminSideMenuItems)
	if (jsonMarshalError != nil) || (jsonMarshalError1 != nil) || (jsonMarshalError0 != nil) {
		log.LoggerSugar.Errorw("menu PutMenuItem json.Marshal Error",
			"module", "application: json.Marshal",
			"error", jsonMarshalError,
			"error0", jsonMarshalError0,
			"error1", jsonMarshalError1,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	_, hmsetError := redis.String(redisConn.Do("HMSET", "menu",
		"menuItems", updatedMenuItems,
		"sideMenuItems", updatedSideMenuItems,
		"adminSideMenuItems", updatedAdminSideMenuItems,
	))
	if hmsetError != nil {
		log.LoggerSugar.Errorw("menu PutMenuItem HMSET menu Error",
			"module", "redis",
			"error", hmsetError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	updateMenuItem := bson.M{
		"$set": bson.M{
			"menuItems." + menuItemName: menuItem,
		},
	}

	updateMenuItemError := Db.C("menu").UpdateId(bson.ObjectIdHex(hMgetreply[0].(string)), updateMenuItem)
	if updateMenuItemError != nil {
		log.LoggerSugar.Errorw("menu PutMenuItem mongo UpdateId Error",
			"module", "mongo",
			"error", updateMenuItemError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}
	log.LoggerSugar.Infow("menu PutMenuItem mongo Success",
		"info", "update MenuItem in menu success.",
	)
	ctx.JSON(http.StatusOK, gin.H{
		"sideMenuItems":      cachedSideMenuItems,
		"menuItems":          cachedMenuItems,
		"adminSideMenuItems": cachedAdminSideMenuItems,
	})
}

// invoke this api means it's the first time to load a menu.
func (this MenuApi) PostMenu(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	redisConn := cache.GlobalUserRedisPool.Get()
	defer Db.Close()
	defer redisConn.Close()

	var (
		sideMenuItems      = make(models.SideMenuItems)
		menu               models.Menu
		adminSideMenuItems = make(map[string]map[string]string)
		menuItem           = make(map[string]models.MenuItem)
	)

	if err := ctx.BindJSON(&menu); err != nil {
		log.LoggerSugar.Errorw("menu PostMenu BindJSON Error",
			"module", "application: BindJson",
			"error", err,
		)
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
		return
	}

	menu.Id = bson.NewObjectId()
	menuItem = menu.MenuItems
	for menuItemName, menuItem := range menu.MenuItems {
		sideMenuItems[menuItemName] = menuItem.SideMenuItems
		adminSideMenuItems[menuItemName] = menuItem.AdminSideMenuItems
	}

	cachedMenuItems, jsonMarshalError := json.Marshal(menuItem)
	cachedSideMenuItems, jsonMarshalError0 := json.Marshal(sideMenuItems)
	cachedAdminSideMenuItems, jsonMarshalError1 := json.Marshal(adminSideMenuItems)
	if (jsonMarshalError != nil) || (jsonMarshalError0 != nil) || (jsonMarshalError1 != nil) {
		log.LoggerSugar.Errorw("menu PostMenu json.Marshal Error",
			"module", "application: json.Marshal",
			"error", jsonMarshalError,
			"error0", jsonMarshalError0,
			"error1", jsonMarshalError1,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	reply, hmsetError := redis.String(redisConn.Do("HMSET", "menu",
		"menuID", menu.Id.Hex(),
		"menuItems", cachedMenuItems,
		"sideMenuItems", cachedSideMenuItems,
		"adminSideMenuItems", cachedAdminSideMenuItems,
	))
	if hmsetError != nil {
		log.LoggerSugar.Errorw("menu PostMenu HMSET menu Error",
			"module", "redis",
			"error", hmsetError,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	} else {
		log.LoggerSugar.Info(menuItem)
		log.LoggerSugar.Infow("menu PostMenu HMSET menu succeed",
			"reply", reply,
		)
	}

	if err := Db.C("menu").Insert(&menu); err != nil {
		log.LoggerSugar.Errorw("PostMenu Insert to Mongo Error",
			"module", "mongo",
			"error", err,
		)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	} else {
		log.LoggerSugar.Info("menu PostMenu Insert to Mongo succeed")
		ctx.JSON(http.StatusCreated, menu)
	}
}
