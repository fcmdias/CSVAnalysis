import { createSignal, onMount } from 'solid-js';
import {     
  Chart,
  LineController,
  CategoryScale,
  PointElement,
  LineElement,
  LinearScale, 
}  from 'chart.js';
import { DefaultChart } from 'solid-chartjs'


function App() {  
  const [chartData, setChartData] = createSignal({ datasets: [] });


  const fetchData = async (sortOrder = 'desc') => {
    try {
      const response = await fetch(`http://localhost:8080/popular?sort=${sortOrder}`);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const electricCars = await response.json();

      const data = {
        labels: electricCars.map(car => `${car.make} ${car.model}`),
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

  const sortAscending = () => fetchData('asc');
  const sortDescending = () => fetchData('desc');

  return (
    <div>
      <h1>Electric Cars by Make and Model</h1>
      <button onClick={sortAscending}>Sort Ascending</button>
      <button onClick={sortDescending}>Sort Descending</button>
      <DefaultChart type="bar" data={chartData()} />
    </div>
  );
}

export default App;
