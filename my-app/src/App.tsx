import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import UserList from "./views/user/UserList";
import UserForm from "./views/user/UserForm";
import GenreList from "./views/genre/GenreList";
import GenreForm from "./views/genre/GenreForm";


const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        {/* <Route path="/" element={<Home />} /> */}
        <Route path="/users" element={<UserList />} />
        <Route path="/users/new" element={<UserForm />} />
        <Route path="/genres" element={<GenreList />} />
        <Route path="/genres/new" element={<GenreForm />} />
      </Routes>
    </Router>
  );
};

export default App;
