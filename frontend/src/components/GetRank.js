const GetRank = () => {
  return (
    <div className='container mt-2'>
      <h2 className="text-center">Get your rank</h2>
      <hr />
      <form className="p-4">
        <div className="form-group">
          <label htmlFor="username" className='form-label'>Username</label>
          <input
            type="text"
            className="form-control"
            id="username"
            placeholder="Enter your username"
          />
        </div>
        <button type="submit" className="btn btn-primary btn-sm mt-3">
          Get Rank
        </button>
      </form>
      <hr />
    </div>
  );
};

export default GetRank;
