import React from "react";

interface GradeSelectorProps {
    grades: Array<{name: string}>;
}

const GradeSelector: React.FC<GradeSelectorProps> = ({grades}) => {
    return (
        <div>
            <h2>Select a Grade:</h2>
            <ul>
                {grades.map((grade, index) => (
                    <li key={index}>{grade.name}</li>
                ))}
            </ul>
        </div>
    );
};

export default GradeSelector;
