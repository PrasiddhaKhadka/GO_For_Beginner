# The 7 Methods — RESTful Resource Convention

This is the classic CRUD over HTTP pattern, same as Rails resources and Laravel resource controllers.  
Most web frameworks map URLs to these 7 actions.

| Method | HTTP Verb   | Route                 | Purpose                        |
| ------ | ----------- | --------------------- | ------------------------------ |
| Index  | GET         | `/campaigns`          | List all campaigns             |
| Create | GET         | `/campaigns/new`      | Show the form to create        |
| Store  | POST        | `/campaigns`          | Save the new campaign to DB    |
| Show   | GET         | `/campaigns/:id`      | Display one campaign           |
| Edit   | GET         | `/campaigns/:id/edit` | Show the form to edit          |
| Update | PUT / PATCH | `/campaigns/:id`      | Save the edited campaign to DB |
| Delete | DELETE      | `/campaigns/:id`      | Remove the campaign            |
