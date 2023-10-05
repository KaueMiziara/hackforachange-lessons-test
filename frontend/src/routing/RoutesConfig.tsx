import React from "react";
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import MainScreen from "../components/MainScreen";
import SubjectSelection from "../components/SubjectSelection";

const RoutesConfig: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<MainScreen />} />
                <Route path="/subjects/:gradeID" element={<SubjectSelection />}/>
            </Routes>
        </Router>
    )
}

export default RoutesConfig;
