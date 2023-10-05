import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

interface LessonExercise {
  id: number;
  title: string;
  content: string;
}

const LessonsAndExercisesScreen: React.FC = () => {
  const {gradeID} = useParams<{gradeID: string}>();
  const {subjectID} = useParams<{subjectID: string}>();

    const [lessons, setLessons] = useState<LessonExercise[]>([]);
    const [exercises, setExercises] = useState<LessonExercise[]>([]);

  useEffect(() => {
    fetch(`http://localhost:8080/api/lessons/${gradeID}/${subjectID}`)
      .then((response) => response.json())
      .then((data) => setLessons(data))
      .catch((error) => console.error("Error fetching lessons: ", error));

    fetch(`http://localhost:8080/api/exercises/${gradeID}/${subjectID}`)
      .then((response) => response.json())
      .then((data) => setExercises(data))
      .catch((error) => console.error("Error fetching exercises: ", error));
  }, [gradeID, subjectID]);

  const combinedData: LessonExercise[] = [...lessons, ...exercises];

  return (
    <div>
      <h2>Lessons and Exercises</h2>
      <ul>
        {combinedData.map((item) => (
          <li key={item.id}>
            <h3>{item.title}</h3>
            <p>{item.content}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default LessonsAndExercisesScreen;
