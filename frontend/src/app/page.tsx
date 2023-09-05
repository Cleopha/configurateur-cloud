'use client';

import { Button, Typography } from '@mui/material';
import Box from '@mui/material/Box';
import React from 'react';

import { usePanelContext } from '../contexts/PanelContext';
import { EPanel, panels } from '../types/panel';

export default function Home() {
	const { indexPanel } = usePanelContext();
	const panelsLength = Object.keys(panels).length;

	return (
		<main>
			<Box display={'flex'} flex={'1'} flexDirection={'column'} sx={{ minHeight: '98vh', minWidth: '100wh' }}>
				<Box
					display={'flex'}
					flex={'1'}
					mb={'12px'}
					borderRadius={'16px'}
					sx={{ boxShadow: 3 }}
					minHeight={'100px'}
					bgcolor={'white'}
				></Box>

				<Box
					display={'flex'}
					flex={'4'}
					borderRadius={'16px'}
					sx={{ border: 1, borderColor: 'lightgray', boxShadow: 3 }}
					bgcolor={'white'}
				>
					<Box display={'flex'} flex={'1'}>
						{/* TODO: Stepper Box */}
						<Box
							display={'flex'}
							flex={'1'}
							flexDirection={'column'}
							sx={{ borderRight: 1, borderColor: 'lightgray' }}
						>
							<Box display={'flex'} flex={'4'} flexDirection={'column'} m={'12px'}>
								{Object.values(panels).map((panel, index) => (
									<Box
										key={index}
										display={'flex'}
										onClick={() => {
											console.log('Click on: ', index);
											indexPanel.setValue?.(index);
										}}
									>
										<Box display={'flex'} flex={'1'} flexDirection={'column'} m={'12px'}>
											<Typography variant="h5" fontSize={'17px'}>
												{panel.title}
											</Typography>
											<Typography fontSize={'12px'} color={'#737376'}>
												0 élément sélectionné
											</Typography>
										</Box>

										{/* TODO: Circle and connector */}
										<Box>
											<Box
												display={'flex'}
												sx={{ cursor: 'pointer' }}
												height={'50px'}
												width={'50px'}
												bgcolor={index === indexPanel.value ? 'beige' : 'lightgray'}
												borderRadius={'100px'}
												alignItems={'center'}
												justifyContent={'center'}
											>
												{React.createElement(panel.icon)}
											</Box>
											{index === panelsLength - 1 ? (
												<></>
											) : (
												<Box display={'flex'} alignItems={'center'} justifyContent={'center'}>
													<Box
														display={'block'}
														height={'60px'}
														sx={{ border: '1px dashed', borderColor: 'lightgray' }}
													></Box>
												</Box>
											)}
										</Box>
									</Box>
								))}
							</Box>

							{/* TODO: Total Box */}
							<Box
								display={'flex'}
								flex={'1'}
								sx={{ borderTop: 1, borderColor: 'lightgray' }}
								p={'12px'}
								minHeight={'100px'}
							>
								<Box
									display={'flex'}
									flex={'1'}
									bgcolor={'beige'}
									borderRadius={'20px'}
									justifyContent={'center'}
									alignItems={'center'}
									flexDirection={'column'}
								>
									<Box>
										<Typography fontSize={'14px'}>Total</Typography>
										<Box display={'flex'} flexDirection={'row'} alignItems={'center'}>
											<Typography fontSize={'22px'} fontWeight={'bold'} marginRight={'8px'}>
												0.00
											</Typography>
											<Typography fontSize={'16px'}>€ / mois</Typography>
										</Box>
									</Box>
								</Box>
							</Box>
						</Box>

						{/* Container Box */}
						<Box display={'flex'} flex={'3'} flexDirection={'column'}>
							{/* TODO: Panel adding */}
							<Box display={'flex'} flex={'10'}>
								{React.createElement(panels[indexPanel.value as EPanel].component, {
									index: indexPanel.value,
								})}
							</Box>

							{/* Next button to stepper */}
							<Box
								display={'flex'}
								flex={'1'}
								sx={{ borderTop: 1, borderColor: 'lightgray' }}
								justifyContent={'flex-end'}
								alignItems={'center'}
							>
								<Box mr={'12px'}>
									<Button
										variant={'contained'}
										sx={{
											backgroundColor: 'beige',
											color: 'black',
											'&:hover': { backgroundColor: '#e9e9d1' },
										}}
										onClick={() => {
											indexPanel.setValue?.((indexPanel.value + 1) % panelsLength);
										}}
									>
										Etape suivante
									</Button>
								</Box>
							</Box>
						</Box>
					</Box>
				</Box>
			</Box>
		</main>
	);
}
