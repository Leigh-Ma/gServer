Client will send login request to gateway. If is a new client, gateway will generate a new user_id, and a new user, and 
    chose a server for this client to login.

When server receives a login request, it will first find the player in it's memory cache, then try load from database,
     then create new (Id is given by gateway). 
 
When client send a search room request, a room should be chosen for the client, and a login room request should be send
     to the server where the room exists, so, client will then "login" the server of the room. All operations after will
     be forward to this server after.

So, a client is "login" to multiply server, but operations only on one server at a time.