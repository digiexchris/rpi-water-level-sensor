import * as React from 'react';

export default function Cell({
                                 content,
                                 header,
                             }) {
    const cellMarkup = header ? (
        <th className="Cell Cell-header">
            {content}
        </th>
    ) : (
        <td className="Cell">
                <div className={content?"sensor sensor-on":"sensor sensor-off"}>
                </div>
        </td>
    );

    return (cellMarkup);
}