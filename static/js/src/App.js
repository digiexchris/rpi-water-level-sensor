import React, {Component} from 'react';
import './App.css';
import DataTable from "./table";

class App extends Component {

    constructor() {
        super();
        this.state = {
            rows:
                {
                    Readings: [
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
                            false
                        ],
                        [
                            false
                        ],
                    ]
                }
        }
    }

    getNewData() {
        fetch('http://localhost:8080/api/readings.json')
            .then(
                results => {
                    return results.json()
                }
            ).then(
            data => {
                //data.reverse();
                this.setState({rows: data.Readings})
            }
        )
    }

    componentDidMount() {
        this.getNewData();

        this.timer = setInterval(()=> this.getNewData(), 5000);
    }

    componentWillUnmount() {
        this.timer = null; // here...
        clearInterval(this.timer)
    }

    render() {
        const headings = [
            'Water Level Sensor',
        ];

        return (
            <div className="App">
                <DataTable headings={headings} rows={this.state.rows}/>
            </div>
        );
    }
}

export default App;

//export default App;
