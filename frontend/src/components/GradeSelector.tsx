import React, { useEffect, useState } from "react";
import ListItem from "./ListItem";
import "./GradeSelector.css";

interface Grade {
    id: number;
    name: string;
}

const GradeSelector: React.FC = () => {
    const [grades, setGrades] = useState<Grade[]>([]);

    useEffect(() => {
        fetch("http://localhost:8080/api/grades")
            .then((response) => response.json())
            .then((data) => setGrades(data))
            .catch((error) => console.error("Error fetching grades: ", error))
    }, []);

    return (
        <div >
            <h2>Select a Grade</h2>
            <div className="grades-container">
                {grades.map((grade) => (
                    <ListItem key={grade.id} name={grade.name} />
                ))}
            </div>
        </div>
    );
};

export default GradeSelector;
