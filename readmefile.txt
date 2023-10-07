# GoLang Blog Platform Assignment

This is a simple blogging platform built using GoLang, Gin framework, and MySQL database. 
It provides basic following operations for managing blog posts.

- Retrieve a list of all blog posts.
- Retieve a specific blog post by ID.
- Create a new blog post.
- Update an existing blog post.
- Delete a blog post.-

## Prerequisites

Before running this application, make sure you have the following:

- GoLang installed on your system
- MySQL database server running
- Postman for testing the APIs (see Postman collection in postmanCollection.txt)

## Setup

1. Clone this repository:

   ```shell
   git clone https://github.com/PATILSHUBHAM69/bloggingPlatform_Assignment.git

2. Create a .env file in the project root directory with the following content, 
replacing the placeholders with your database credentials:
DB_USERNAME=your_db_username
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

3. Install the required Go packages:

-- go mod tidy

4. Run the application:

-- go run main.go


5. Check all api on postman mention in postmanCollection.txt file

-- POST /postNewBlog: Create a new blog post.
-- GET /getAllBlogsposts: Get all blog posts.
-- GET /getBlogpostByID/{:userID}: Get a blog post by its ID.
-- PUT /updateBlogpostByID: Update a blog post by its ID.
-- DELETE /deleteBlogpostByID: Delete a blog post by its ID.