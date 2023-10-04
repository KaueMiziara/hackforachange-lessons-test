import GradeSelector from "./components/GradeSelector";
import data from "./data/database.json";
import "./App.css"

function App() {
  return (
    <div>
      <GradeSelector grades={data.grades} />
    </div>
  );
}

export default App;
