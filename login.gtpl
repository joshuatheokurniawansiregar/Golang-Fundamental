<html>
    <head>
        <title>hello</title>
    </head>
    <body style="background-color:white;">
        <form action="/login" method="POST">
             Username:<input type="text" name="username">
             Password:<input type="password" name="password">
             Age: <input type="text" name="age">
             <input type="hidden" name="token" value="{{.}}">
            <input type="submit" value="Login"/>
        </form>
    </body>