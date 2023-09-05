import React from 'react';

interface IPanel {
	index: number;
}

export default function Panel({ index }: IPanel) {
	return <div>{index}</div>;
}
