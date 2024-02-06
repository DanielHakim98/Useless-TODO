function TodoItem() {
  return (
    <div>
      <div>
        <h1>TODO Title</h1>
        <span>TODO Description</span>
      </div>
      <div>
        <button onClick={(e) => console.log(e)}>Complete</button>
        <button onClick={(e) => console.log(e)}>Delete</button>
      </div>
    </div>
  );
}
export default TodoItem;
