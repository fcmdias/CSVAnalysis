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
  const [selectedFilter, setSelectedFilter] = createSignal('all');




  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/byyear?filter=${selectedFilter()}`);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const electricCars = await response.json();

      const data = {
        labels: electricCars.map(car => `${car.year}`),
        datasets: [{
          label: 'Electric Cars',
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

  return (
    <div className="container mt-4">
      <h1>By Year</h1>
      <p>This dataset shows the Battery Electric Vehicles (BEVs) and Plug-in Hybrid Electric Vehicles (PHEVs) that are currently registered through Washington State Department of Licensing (DOL).</p>
      <div className="btn-group mb-3">
        <button className={`btn ${selectedFilter() === 'all' ? 'btn-primary' : 'btn-outline-primary'}`} onClick={applyFilter('all')}>All</button>
        <button className={`btn ${selectedFilter() === 'electric' ? 'btn-primary' : 'btn-outline-primary'}`} onClick={applyFilter('electric')}>Electric</button>
        <button className={`btn ${selectedFilter() === 'hybrid' ? 'btn-primary' : 'btn-outline-primary'}`} onClick={applyFilter('hybrid')}>Hybrid</button>
      </div>
      <DefaultChart type="bar" data={chartData()} />
    </div>
  );
}

export default ByYear;
