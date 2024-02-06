function TodoItem() {
  return (
    <div className="Card">
      <div className="Card--text">
        <h1>Insert TODO title here</h1>
        <span></span>
      </div>
      <div className="Card--button">
        <button onClick={(e) => console.log(e)}>Complete</button>
        <button
          onClick={(e) => console.log(e)}
          className="Card--button__delete"
        >
          Delete
        </button>
      </div>
    </div>
  );
}
export default TodoItem;
