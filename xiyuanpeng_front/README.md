Hi, In the "embedding all comments in one document" part, for threaded discussions, you mentioned
that "To reply to a comment, the following code assumes that it can retrieve the ‘path’ as a list of positions",
which is really a clever idea, but the thing is I can't find a practical way to get the path list.

I've considered maintaining a field named internalPath which demonstrate the path of the comment in the whole tree,
but it seems impractical to do so because it's very hard for mongodb to find each comment's index in the arrays given
 the depth of the whole tree.

And I've considered maintaining the path in frontend, which is not as hard to implement as doing it in database level,
 but that will certainly decrease the reliability, right?

So my question is how to retrieve the 'path' list.

Looking forward your reply, thanks.
