import BalanceOutlinedIcon from '@mui/icons-material/BalanceOutlined';
import DesktopWindowsOutlinedIcon from '@mui/icons-material/DesktopWindowsOutlined';
import TravelExploreOutlinedIcon from '@mui/icons-material/TravelExploreOutlined';
import ViewInArOutlinedIcon from '@mui/icons-material/ViewInArOutlined';
import WidgetsOutlinedIcon from '@mui/icons-material/WidgetsOutlined';
import React from 'react';

import Panel from '../components/Panel/Panel';

export enum EPanel {
	INSTANCE,
	BLOCK_STORAGE,
	OBJECT_STORAGE,
	IP_ADDRESS,
	LOAD_BALANCER,
}

interface IPanelDescription {
	title: string;
	description: string;
	icon: React.ComponentType;
	component: React.ComponentType<any>;
}

export const panels: { [key in EPanel]: IPanelDescription } = {
	[EPanel.INSTANCE]: {
		title: 'Instance',
		description: 'Serveurs Cloud évolutifs et à la demande.',
		icon: DesktopWindowsOutlinedIcon,
		component: Panel,
	},
	[EPanel.BLOCK_STORAGE]: {
		title: 'Block Storage',
		description:
			'Block storage haute performance (SSD PCIe NVMe dernière génération) et haute disponibilité (3 répliques) propulsé par Ceph.',
		icon: WidgetsOutlinedIcon,
		component: Panel,
	},
	[EPanel.OBJECT_STORAGE]: {
		title: 'Object Storage',
		description:
			"Un stockage d'objets Swift compatible S3 pour stocker et accéder à vos ressources, fichiers et métadonnées.",
		icon: ViewInArOutlinedIcon,
		component: Panel,
	},
	[EPanel.IP_ADDRESS]: {
		title: 'Adresses IP publiques\n',
		description: '',
		icon: TravelExploreOutlinedIcon,
		component: Panel,
	},
	[EPanel.LOAD_BALANCER]: {
		title: 'Load balancer',
		description:
			'Un load balancer (répartiteur de charge) a pour tâche de répartir la charge de travail entre un ou plusieurs serveurs.',
		icon: BalanceOutlinedIcon,
		component: Panel,
	},
};
