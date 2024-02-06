import "./App.css";
import TodoItem from "./Todo/TodoItem";

function App() {
  return [
    <TodoItem key={"1"} />,
    <TodoItem key={"2"} />,
    <TodoItem key={"3"} />,
  ];
}

export default App;
