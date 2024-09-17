const SubmitForm = () => {
  return (
    <>
      <div className="container mt-3">
        <div className="text-center">
          <h2>Submit your score !!</h2>
          <hr />
        </div>
        <form className="row gx-3 gy-2">
          <div className="col-md-6">
            <label htmlFor="username" className="form-label">
              Username
            </label>
            <input
              type="text"
              className="form-control"
              id="username"
              placeholder="Enter your username"
            />
          </div>
          <div className="col-md-6">
            <label htmlFor="state" className="form-label">
              State
            </label>
            <input
              type="text"
              className="form-control"
              id="state"
              placeholder="Enter your state"
            />
          </div>
          <div className="col-md-6">
            <label htmlFor="country" className="form-label">
              Country
            </label>
            <input
              type="text"
              className="form-control"
              id="country"
              placeholder="Enter your country"
            />
          </div>
          <div className="col-md-6">
            <label htmlFor="score" className="form-label">
              Score
            </label>
            <input
              type="number"
              className="form-control"
              id="score"
              placeholder="Enter your score"
            />
          </div>
          <div className="col-12">
            <button type="submit" className="btn btn-primary btn-sm">
              Submit
            </button>
          </div>
        </form>
      </div>
    </>
  );
};

export default SubmitForm;
