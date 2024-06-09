# point counter serber
## auth

<details>
 <summary><code>POST</code> <code><b>/auth/user/register/</b></code> <code>create new account for user aka register</code></summary>

##### Parameters

> | name     | type     | data type | description |
> |----------|----------|-----------|-------------|
> | username | required | string    | N/A         |
> | email    | requrred | string    | N/A         |
| | password | required | string    | must be hash|


##### Responses

> | http code     | content-type              | response                                                                                                                                                  |
> |---------------|---------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"user": {"CreatedAt": "<date>","UpdatedAt": "<date","DeletedAt": null,"ID": <number>,"role_id": <number>,"username": "<string>>","email": "<string>"}}' |
> | `400`         | `text/html;charset=utf-8` | {"error":"..."}                                                                                                                                           |

##### example data

```json
{
	"username": "example",
	"email":"user@axample.com",
	"password":"example"
}
```
[//]: # (##### Example cURL)

[//]: # (> ```javascript)

[//]: # (>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8889/)

[//]: # (> ```)

</details>
<details>
<summary><code>POST</code> <code><b>/auth/user/login/</b></code> <code>login into existing account</code></summary>

##### Parameters

> | name     | type     | data type | description |
> |----------|----------|-----------|-------------|
|          | username | reauired | string    | N/A |
|          | password | required | string    | must be hash


##### Responses

> | http code | content-type       | response                                                                                                                                                  |
> |-----------|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
|               | 200       | `application/json` | `{"message":"<string>", "token":<jwttoken>,"username":"<string>"}`
|           | 400       | `application/json`   | `{"error":"string"}`

##### example data

```json
{
	"username": "example",
	"password": "password"
}
```
</details>

## Public

### Room read

<details>
<summary><code>GET</code> <code><b>/api/view/rooms/</b></code><code>get all rooms</code></summary>

##### Params

none

##### response

[//]: # (TODO)

</details>

<details>
<summary><code>GET</code> <code><b>/api/view/room/:id</b></code><code>get room data</code></summary>

##### Params

none

##### response

[//]: # (TODO)

</details>
<details>
<summary><code>GET</code> <code><b>/api/view/room/:id/stats</b></code><code>get all stats for specyfic room</code></summary>

##### Params

none

##### response

[//]: # (TODO)

</details>

### read users
<details>
<summary><code>GET</code><code><b>/api/view/users</b></code><code>get list of all users</code></summary>

##### params
none

##### response

</details>

## protected

### get your data
<details>
<summary><code>GET</code><code><b>/api/me</b></code><code>get your user data</code></summary>

##### response data

> | name     | data type | description |
> |----------|-----------|-------------|
|         | ID       | int       |      
|      | username | string    |
|          | email    | string    |
</details>

### room crud
<details>
<summary><code>POST</code><code><b>/api/room/</b></code><code>create new room</code></summary>

##### Params

> | name    | type           | data type | description |
> |---------|----------------|-----------|-------------|
|          | user_id | requred        | number    | id of user who creates room |
|         | name    | required| string    | unique name of room |

##### response 

array[string]

[//]: # (TODO)
</details>

### post stat

<details>
<summary><code>POST</code><code><b>/api/stat/</b></code><code>post stat to the room</code></summary>

##### Params

| name |type | datatype | description|
|------|----|---------|-----------|
| Value| requred| number| score, point ect|
|Comment| optional|string| some type of comment|

##### Respoinse 

[//]: # (TODO)


</details>