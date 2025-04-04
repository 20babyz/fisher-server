db = db.getSiblingDB("dev-db");

db.createCollection("users");
db.users.insertMany([
    { username: "admin", password: "admin123", role: "admin" },
    { username: "user1", password: "user123", role: "user" }
]);

// 로그 출력
print("Database 'dev-db' and 'users' collection initialized.");