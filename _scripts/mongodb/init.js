db = db.getSiblingDB("bank");

db.createUser({
  user: "dev",
  pwd: "dev",
  roles: [
    {
      role: "root",
      db: "admin",
    },
  ],
});

let accounts = db.createCollection("accounts");
db.accounts.createIndex({ cpf: 1 }, { unique: true });

db.createCollection("transfers");
