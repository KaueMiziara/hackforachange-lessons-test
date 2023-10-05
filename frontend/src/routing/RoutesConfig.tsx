import React from "react";
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import MainScreen from "../components/MainScreen";

const RoutesConfig: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<MainScreen />} />
            </Routes>
        </Router>
    )
}

export default RoutesConfig;
