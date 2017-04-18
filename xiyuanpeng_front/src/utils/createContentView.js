import SideMenuView from '../components/SideMenuView'

// This is a factory function for dynamically creating root-level list views,
// since they share most of the logic except for the type of items to display.
// They are essentially higher order components wrapping ItemList.vue.
export default function createListView (type) {
  return {
    name: `${type}-menu-view`,
    render (h) {
      return h(SideMenuView, { props: { type } })
    }
  }
}
