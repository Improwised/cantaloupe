import {
	FETCH_INCIDENTS_SUCCESS,
	FETCH_INCIDENTS_FAILULER,
	SET_FIELDS,
	RESET_STATE_INCIDENTS,
	FETCH_INCIDENT_INFORMATION_SUCCESS,
	FETCH_INCIDENT_INFORMATION_FAILULER,
	ADD_INCIDENT_UPDATE_SUCCESS,
	ADD_INCIDENT_UPDATE_FAILULER,
	COMPONENT_REPLACED_SUCCESS,
	COMPONENT_REPLACED_FAILULER,
	COMPONENT_ACTIVE_SUCCESS,
	COMPONENT_ACTIVE_FAILULER,
	COMPONENT_DEACTIVE_SUCCESS,
	COMPONENT_DEACTIVE_FAILULER,
} from '../../constants';

const initialState =
	{
		Incidents : [],
		isFetching : true,
		fetched : false,
		err : '',

		description: '',
		title: '',
		recorder: '',
		machine: '',

		resolvedBy: '',
		updatedBy: '',
		serialNo: '',
		component: '',

		active: '',
	}

export default function incidents (state = initialState, action) {
	switch (action.type) {
		case FETCH_INCIDENTS_SUCCESS:
			return Object.assign({}, state, { Incidents : action.response.data, isFetching:false, fetched:true, err:''});

		case FETCH_INCIDENTS_FAILULER:
			return Object.assign({}, state, { Incidents : '', isFetching:false, fetched:false, err:action.err});

		case SET_FIELDS:
			state[action.field] = action.value;
			return Object.assign({}, state, { state });

		case RESET_STATE_INCIDENTS:
			state['description'] = ''
			state['title'] = ''
			state['recorder'] = ''
			state['machine'] = ''
			state['updatedBy'] = ''
			state['resolvedBy'] = ''
			state['serialNo'] = ''
			state['component'] = ''

		case FETCH_INCIDENT_INFORMATION_SUCCESS:
			return Object.assign({}, state, { Incidents : action.response.data, isFetching:false, fetched:true, err:''});

		case FETCH_INCIDENT_INFORMATION_FAILULER:
			return Object.assign({}, state, { Incidents :'', isFetching:false, fetched:true, err:action.err});

		case ADD_INCIDENT_UPDATE_SUCCESS:
			console.log('ADD_INCIDENT_UPDATE_SUCCESS');

		case ADD_INCIDENT_UPDATE_FAILULER:
			console.log('ADD_INCIDENT_UPDATE_FAILULER');

		case ADD_INCIDENT_UPDATE_FAILULER:
			console.log('ADD_INCIDENT_UPDATE_FAILULER');

		case ADD_INCIDENT_UPDATE_FAILULER:
			console.log('ADD_INCIDENT_UPDATE_FAILULER');

		case COMPONENT_ACTIVE_SUCCESS:
			return Object.assign({}, state, { Incidents :'', isFetching:false, fetched:true, err:'', active:'Active'});

		case COMPONENT_ACTIVE_FAILULER:
			return Object.assign({}, state, { Incidents :'', isFetching:false, fetched:true, err:''});

		case COMPONENT_DEACTIVE_SUCCESS:
			return Object.assign({}, state, { Incidents :'', isFetching:false, fetched:true, err:'Error', active:'Deactive'});

		case COMPONENT_DEACTIVE_FAILULER:
			return Object.assign({}, state, { Incidents :'', isFetching:false, fetched:true, err:'Error'});


		default:
			return state
	}
}
