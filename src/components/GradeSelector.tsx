import React from "react";
import ListItem from "./ListItem";

interface GradeSelectorProps {
    grades: Array<{name: string}>;
}

const GradeSelector: React.FC<GradeSelectorProps> = ({grades}) => {
    return (
        <div>
            <h2>Select a Grade:</h2>
            <div>
                {grades.map((grade, index) => (
                    <ListItem key={index} name={grade.name} />
                ))}
            </div>
        </div>
    );
};

export default GradeSelector;
