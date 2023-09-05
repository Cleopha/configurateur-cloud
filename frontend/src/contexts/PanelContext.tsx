'use client';

import React, { createContext, ReactNode, useContext, useState } from 'react';

interface SetStateContext<T> {
	value: T;
	setValue?: React.Dispatch<React.SetStateAction<T>>;
}

interface IPanelContext {
	indexPanel: SetStateContext<number>;
}

const PanelContext = createContext<IPanelContext>({
	indexPanel: {
		value: 0,
	},
});

export const PanelProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
	const [indexPanel, setIndexPanel] = useState<number>(0);

	return (
		<PanelContext.Provider
			value={{
				indexPanel: {
					value: indexPanel,
					setValue: setIndexPanel,
				},
			}}
		>
			{children}
		</PanelContext.Provider>
	);
};

export const usePanelContext = () => useContext(PanelContext);
