namespace go wangxingge.thrift_clientpool.examples.bookservice
namespace csharp wangxingge.thrift_clientpool.examples.bookservice

include "./entity.thrift"

service BookService
{
    entity.Book         GetBookById(1:string bookId)
    entity.Book         GetBookByName(1:string bookName)
    list<entity.Book>   GetAllBooks()
    bool                AddBook(1:entity.Book bookInfo)
    bool                RemoveBook(1:string bookId)
    bool                DefaultKeepAlive(1:string clientId)
}

service UserService
{
    list<entity.Book>   GetUserBooks(1:string userId)
    entity.User         GetUserInfo(1:string userId)
    list<entity.User>   GetAllUserInfo()
    bool                AddUser(1:entity.User userInfo)
    bool                RemoveUser(1:string userId)
    bool                UpdateUserAvatar(1:string userId, 2:binary avatar)
    bool                DefaultKeepAlive(1:string clientId)
}