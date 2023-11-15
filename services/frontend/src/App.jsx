import 'bootstrap/dist/css/bootstrap.min.css';
import ByYear from './components/ByYear/ByYear'
import ByMakeModel from './components/ByMakeModel/ByMakeModel'

function App() {  
  
  return (
    <div className="container mt-4">
      <ByMakeModel/>
      <ByYear/>
    </div>
  );
}

export default App;
