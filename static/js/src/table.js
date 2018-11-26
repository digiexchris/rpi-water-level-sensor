import * as React from 'react';
import Cell from './cell';

import _ from 'lodash'

export default class DataTable extends React.Component {
    renderHeadingRow = (_cell, cellIndex) => {
        const {headings} = this.props;

        return (
            <Cell
                key={`heading-${cellIndex}`}
                content={headings[cellIndex]}
                header={true}
            />
        )
    };

    renderRow = (_row, rowIndex) => {
        // const {rows} = this.props;
        //console.log(_row, rowIndex);
        return (
            <tr key={`row-${rowIndex}`}>
                <Cell
                    key={`cell-${rowIndex}`}
                    content={_row}
                />
            </tr>
        )
    };

    render() {
        const {headings, rows} = this.props;

        this.renderHeadingRow = this.renderHeadingRow.bind(this);
        this.renderRow = this.renderRow.bind(this);

        const theadMarkup = (
            <tr key="heading">
                {headings.map(this.renderHeadingRow)}
            </tr>
        );

        const tbodyMarkup = _.values(rows).reverse().map(this.renderRow);
        return (
            <table className="Table">
                <thead>{theadMarkup}</thead>
                <tbody>{tbodyMarkup}</tbody>
            </table>
        );
    }
}