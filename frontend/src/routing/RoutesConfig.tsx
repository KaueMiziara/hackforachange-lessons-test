import React from "react";
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import MainScreen from "../components/MainScreen";
import LessonsAndExercisesScreen from "../components/LessonsAndExercisesScreen";
import SubjectSelection from "../components/SubjectSelection";

const RoutesConfig: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<MainScreen />} />
                <Route path="/subjects/:gradeID" element={<SubjectSelection />}/>
                <Route
                 path="/lessons-and-exercises/:gradeID/:subjectID" 
                 element={<LessonsAndExercisesScreen />} 
                />
            </Routes>
        </Router>
    )
}

export default RoutesConfig;
