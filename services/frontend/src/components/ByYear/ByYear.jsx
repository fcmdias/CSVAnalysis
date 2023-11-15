import { createSignal, onMount } from 'solid-js';
import {     
  Chart,
  LineController,
  CategoryScale,
  PointElement,
  LineElement,
  LinearScale, 
}  from 'chart.js';
import { DefaultChart } from 'solid-chartjs';
import 'bootstrap/dist/css/bootstrap.min.css';


function ByYear() {  
  const [chartData, setChartData] = createSignal({ datasets: [] });
  const [selectedData, setData] = createSignal([]);
  const [selectedFilter, setSelectedFilter] = createSignal('all');
  const [selectedAIData, setAIData] = createSignal('');
  const [isLoading, setIsLoading] = createSignal(false);

  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/byyear?filter=${selectedFilter()}`);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const electricCars = await response.json();

      setData(electricCars)
      const data = {
        labels: electricCars.map(car => `${car.year}`),
        datasets: [{
          label: 'Electric Cars By Year',
          data: electricCars.map(car => car.total),
          backgroundColor: 'rgba(0, 123, 255, 0.5)',
          borderColor: 'rgba(0, 123, 255, 1)',
          borderWidth: 1
        }]
      };

      setChartData(data);
    } catch (error) {
      console.error('Failed to fetch electric cars:', error);
    }
  };

  onMount(() => {    
    fetchData();
    Chart.register(LineController, CategoryScale, PointElement, LineElement, LinearScale)
  });


  const applyFilter = (filter) => () => {
    setSelectedFilter(filter);
    fetchData()
  };
 
 
  
  const askAI = async () => {
    setIsLoading(true);
    try {
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ data: selectedData() })
      };
      const response = await fetch(`http://localhost:8000/byyear?filter=${selectedFilter()}`, requestOptions);

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      setAIData(data);
      console.log("Received data:", data);

    } catch (error) {
      console.error('Failed to fetch electric cars:', error);
    } finally {
      setIsLoading(false); // Stop loading regardless of outcome
    }
  };
  
  return (
    <div className="container mt-4">
      <h1>Cars Registered Through (DOL) By Year</h1>
      <div className="btn-group mb-3">
        <button className={`btn ${selectedFilter() === 'all' ? 'btn-primary' : 'btn-outline-primary'}`} onClick={applyFilter('all')}>Electric and Hybrid</button>
        <button className={`btn ${selectedFilter() === 'electric' ? 'btn-primary' : 'btn-outline-primary'}`} onClick={applyFilter('electric')}>Electric</button>
        <button className={`btn ${selectedFilter() === 'hybrid' ? 'btn-primary' : 'btn-outline-primary'}`} onClick={applyFilter('hybrid')}>Hybrid</button>
      </div>
      <DefaultChart type="bar" data={chartData()} />
      
      <button 
        className={`btn ${isLoading() ? 'btn-secondary' : 'btn-primary'}`} 
        onClick={askAI}
        disabled={isLoading()}
      >
        {isLoading() ? 'Loading...' : 'Ask AI'}
      </button>
      {isLoading() && 
        <div className="spinner-border text-primary" role="status">
          <span className="sr-only">Loading...</span>
        </div>
      }
      <p>{selectedAIData()}</p>
    </div>
  );
}

export default ByYear;
