import './App.css';
import Login from './features/auth/login/login';


function App() {
  const handleLogin = (credentials) => {
      console.log('Login credentials:', credentials);
      // Add your login logic here
  };

  return (
      <div className="App">
          <Login onLogin={handleLogin} />
      </div>
  );
}

export default App;
