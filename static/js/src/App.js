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
                },
            Error: false
        }
    }

    getNewData() {
        fetch('/api/readings.json')
            .then(
                results => {
                    return results.json()
                }
            ).then(
            data => {
                //data.reverse();
                this.setState({rows: data.Readings, Error: data.Error})
            }
        )
            .catch((e) => {
            this.setState({Error: e.toString()});
        })
    }

    componentDidMount() {
        this.getNewData();

        this.timer = setInterval(()=> this.getNewData(), 5000);
    }

    componentWillUnmount() {
        this.timer = null; // here...
        clearInterval(this.timer)
    }

    renderErrorWarnings () {
        var errors = this.state.Error;

        if(errors !== "false") {
            return (
                <p className="error">An error has occurred. You should not trust these results. {errors}</p>
            )
        } else {
            return (
                <p>

                </p>
            )
        }


    };

    render() {
        const headings = [
            'Water Level Sensor',
        ];

        const renderedErrors = this.renderErrorWarnings()

        return (
            <div className="App">
                {renderedErrors}
                <DataTable headings={headings} rows={this.state.rows}/>
            </div>
        );
    }
}

export default App;

//export default App;
