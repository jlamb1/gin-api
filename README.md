# How to run
`bazel run //:gin-api` - and that's it! available on localhost:3000

if you do not have bazel, simply run `./bazelisk-darwin build //...`

## Available Routes

**/ping** - GET
 - returns `pong`

**/user/:name** -GET
 - returns "hello {name}"
 - for example, /user/john returns "hello john"

**/form** - POST
 - submit keys:
   
   user: string
   
   message: string

   curl example:

```bash
curl --request POST \
  --url http://localhost:3000/form \
  --header 'Content-Type: multipart/form-data' \
  --form user=john \
  --form 'message=hi there'
  ```

returns:
```json
{
  "message": "hi there",
  "status": "posted",
  "user": "john"
}
```
