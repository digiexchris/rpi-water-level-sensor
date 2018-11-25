import React, { Component } from 'react';
import './App.css';
import DataTable from "./table";

class App extends Component {
  render() {
      const headings = [
          'Water Level Sensor',
      ];

      const rows = [
          [
              false
          ],
          [
              false
          ],
          [
              false
          ],
          [
              true
          ],
          [
              true
          ],
      ];

    return (
      <div className="App">
          <DataTable headings={headings} rows={rows} />
      </div>
    );
  }
}

export default App;
