<h1>How to Run the Application</h1>
<h2>Prerequisites</h2>
Before running the application, ensure you have the following installed:
<br><br/>
<ul>
    <li>Go (latest version recommended)</li>
    <li>Docker Compose</li>
    <li>PostgreSQL database</li>
</ul>
<h2>Setup</h2>
<ol>
  <li>Clone the repository:<br><br/><img src="https://github.com/user-attachments/assets/5ae1ce3f-4969-442b-95d1-79367149f21b"/></li>
  <li>Create .env file with your database credentials.</li>
</ol>
<h2>Running the Application</h2>
<h3>Using Docker</h3>
  <img src="https://github.com/user-attachments/assets/f22380bf-61f0-451a-8fdf-e8353a2e9d89"/>
  <br></br>This will start the PostgreSQL database and the application server.
<h3>Running Locally</h3>
  <ol>
    <li>Start the PostgreSQL database and ensure it matches the .env configuration.</li>
    <li>Initialize the database:<br></br><img src="https://github.com/user-attachments/assets/eca8e5c9-70a0-4242-b90a-b32d3c1ac4f1"/><br></br></li>
    <li>The server should be running at:<br></br><img src="https://github.com/user-attachments/assets/b63a4b6f-fefc-494c-bcd3-245ba02e74b9"/></li>
  </ol>
<h1>GraphQL Queries and Mutations</h1>
<h2>Add a Word Translation</h2>
<h3>Mutation</h3>
<img src="https://github.com/user-attachments/assets/d04f2d3f-fec4-4c59-8aea-5b0b471fef81"/>
<h3>Response Example</h3>
<img src="https://github.com/user-attachments/assets/8e0080d6-4f00-4860-8389-974cdd87287a"/>
<h2>Get Translation for a Word</h2>
<h3>Query</h3>
<img src="https://github.com/user-attachments/assets/b0162df5-866b-4c29-aace-ab5ca694265c"/>
<h3>Response Example</h3>
<img src="https://github.com/user-attachments/assets/96e26aa8-3f4e-47c7-a43d-0d452196a0e9"/>
<h3>Query</h3>
<img src="https://github.com/user-attachments/assets/1ffebc32-bcbc-402b-b589-d865caa42305"/>
<h3>Response Example</h3>
<img src="https://github.com/user-attachments/assets/17ef97cd-c62c-4fd1-ad75-7fd28e1d59ca"/><br></br>
Now you can use GraphQL clients like Postman to interact with your API.

