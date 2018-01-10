package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "User List",
        "GET",
        "/users",
        GetUsers,
    },
    Route{
        "Create User",
        "POST",
        "/users",
        CreateUser,
    },
    Route{
        "Get User Relations",
        "GET",
        "/users/{user_id}/relationships",
        GetUserRelations,
    },
    Route{
        "Create/update relationship state to another user",
        "PUT",
        "/users/{user_id}/relationships/{other_user_id}",
        SetUserRelation,
    },
}
