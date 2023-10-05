# Hack for a Change - Team Swift

## Run backend
```bash
  $ cd backend
  $ go run .
```

The server will open on ```localhost:8080``` <br>
Endpoints:
- GET grades: ```/api/grades```
- GET subjects by grade ID: ```/api/subjects/{gradeID}```
- GET lessons by subject and grade ID: ```/api/lessons/{gradeID}/{subjectID}```
- GET exercises by subject and grade ID: ```/api/exercises/{gradeID}/{subjectID}```

## Run frontend
```bash
  $ cd frontend
  $ npm start
```
The server will open on ```localhost:3000``` <br>
