# Backend
- An API for a dating app.
### Get started
- Run `docker-compose up`
- This should create a new mysql db, run the seed, and start the API running at http://localhost:8000

### Schema overview
![Screenshot](/images/schema.png)
- The approach for the schema for this PoC was to keep things simple with just two tables, one for all users and another to contain all of their swipes with a preference (YES/NO).
#### Future improvements
- At the moment we are returning true for a match when there is a pre-existing swipe that has the preference YES for the other user. However, in the future you can imagine a "matches" table where a record would be created in this instance.
- The other future improvement is that the preference field on swipes should probably be an enum rather than a string, and maybe gender.
- We should probably filter out the profiles that have already been matched in the `/profiles` endpoint but for now it returns all of them except the current user as it is easier to test like this with our limited seed data.

### Architecture overview
The API is structured into the following main packages:

### `api`
#### `controllers`
This package contains the controllers for the API. A controller is responsible for handling HTTP requests and returning appropriate responses. Each controller is responsible for a specific set of routes and actions.

#### `models`
This package contains the data models for the API. A model represents a data entity in the system, such as a user or a swipe. The models are typically used by the controllers to perform database operations.

#### `middleware`
This package contains middleware functions that are executed before and after each request. At the moment we only have a middleware to authenticate the jwt.

#### `auth`
This package contains authentication-related functionality, such as functions for generating and validating JSON Web Tokens (JWTs).

#### `utils`
This package contains a utility function to format errors.

### `config`
In addition to the packages in the `api` directory, there is also the `config` package that contains the code that is run when the app starts. When the API starts from `main.go`, it loads the environment variables, initializes the connection to the database, runs migrations, seeds the database, initializes the router and runs the router.

#### Future improvements
- In the future as the application grows it may be necessary to refactor the architecture so that we introduce a `service` layer between the `controllers` and the `models`. In this pattern we would manage all of the validation of our HTTP requests at the controller layer (e.g. ensuring each request has the correct fields in the body), and then move the business logic into our service which would then interact with our models.
- Another area of improvement is how we handle errors, the `FormatError` function is fairly basic at the moment so we could extend + standardize this to ensure more consistency across our API. Maybe move it into a middleware.

# Other useful commands
To force the DB to reset -> `docker-compose up --force-recreate db`
