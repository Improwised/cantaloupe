import React, {Component} from 'react'
import { textarea, Field, FieldArray, reduxForm } from 'redux-form'
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import Select from 'react-select';
import { HelpBlock } from 'react-bootstrap'
import 'react-select/dist/react-select.css';
import './styles.css';

class Add_incident extends Component {
	constructor(props) {
		super(props);
		this.state = {
			machine : '',
			component : '',

			recorderErr: '',
			titleErr: '',
			componentErr: ''
		};

		this.handleSubmit = this.handleSubmit.bind(this);
		this.handleFields = this.handleFields.bind(this);
		this.handleMachinesChange = this.handleMachinesChange.bind(this);
		this.handleComponentsChange = this.handleComponentsChange.bind(this);
	}

	componentDidMount() {
		this.props.actions.fetchMachines(false, this.props.dispatch);
		this.props.actions.fetchComponents(false, this.props.dispatch);
	}

	handleSubmit(e) {
		e.preventDefault();
		let componentId = this.state.component;
		let machine = this.state.machine;
		let recorder = this.props.props.incidents.recorder;
		let title = this.props.props.incidents.title;
		let description = this.props.props.incidents.description;

		this.setState({recorderErr: ''});
		this.setState({titleErr: ''});
		this.setState({componentErr: ''});
		let status = true;

		if(recorder == '') {
			this.setState({recorderErr: '*Required'});
			status = false;
		}
		if(componentId == '') {
			this.setState({componentErr: '*Required'})
			status = false;
		}
		if(title == '') {
			this.setState({titleErr: '*Required'})
			status = false;
		}

		if(status) {
			this.props.actions.addIncident(componentId , recorder, title, description, this.props.dispatch)
		}
	}

	handleMachinesChange(event) {
		this.setState({machine: event.value});
		this.props.actions.fetchMachineInformation(event.value)
	}

	handleComponentsChange(event) {
		this.setState({component: event.value});
	}

	handleFields(event) {
		this.props.actions.setFieldValue(event.target.id, event.target.value);
	}

	render() {
	//======================== style =============================================

		let borderStyle = {
			border: 'solid',
			borderWidth: '2px',
			padding: '20px 25px 20px 100px'
		};

	//============================================================================

		let Machines = [];
		for(let i = 0; i < this.props.props.machines.Machines.length; i++){
			let MachineInfo = { value : this.props.props.machines.Machines[i].Id, label: this.props.props.machines.Machines[i].Name };
			Machines[i] = MachineInfo;
		}

		let Components = [];
		//====== will gives only those components which are connected to perticuller machine ...
		if(this.props.props.machines.MachineInfo.Components) {
			for(let i = 0; i < this.props.props.machines.MachineInfo.Components.length; i++) {
				let ComponentInfo = { value : this.props.props.machines.MachineInfo.Components[i].Id, label: this.props.props.machines.MachineInfo.Components[i].Name };
				Components[i] = ComponentInfo;
			}
		}
		//==========================================================================

		//== will gives all the components regardless connected with any machine ===
		//== By default it will shows all the components ===========================
		else {
			for(let i = 0; i < this.props.props.components.Components.length; i++) {
				let ComponentInfo = { value : this.props.props.components.Components[i].Id, label: this.props.props.components.Components[i].Name };
				Components[i] = ComponentInfo;
			}
		}
		//==========================================================================

		const { handleSubmit, pristine, reset, submitting } = this.props
		return (
			<div>
				<h2 className="center"> Record New Incident </h2>
				<div style={borderStyle}>
					<div className="clearfix form-group">

						<div className = "col-lg-2 col-lg-offset-2">
							<label >Recorder*</label>
							<input
								className="textboxSize"
								id="recorder"
								name="Recorder"
								onChange={ this.handleFields }
								placeholder="Recorder"
								type="text"
								value={ this.props.props.incidents.recorder }
							/>
							<HelpBlock className="errFontStyle"> {this.state.recorderErr} </HelpBlock>

					</div>
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Machine ( optional )</label>
							<Select
							 	clearable={false}
							 	id="machine"
							 	name="Machine"
							 	options={ Machines }
						 		onChange={ this.handleMachinesChange }
								value={ this.state.machine }
							/>
						</div>
					</div>

					<div className="clearfix form-group">
						<div className="col-lg-2 col-lg-offset-2">
							<label >Title*</label>

							<input
								className="textboxSize"
								name="Title" id="title"
								onChange={ this.handleFields }
								placeholder="Title"
								type="text"
								value={ this.props.props.incidents.title }
						 	/>

							<HelpBlock className="errFontStyle"> {this.state.titleErr} </HelpBlock>
						</div>
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Component*</label>
							<Select
								clearable={false}
								name="form-field-name"
								options={ Components }
								onChange={ this.handleComponentsChange }
								value={ this.state.component }
							/>
							<HelpBlock className="errFontStyle"> {this.state.componentErr} </HelpBlock>
						</div>
					</div>

					<div className="clearfix form-group">
						<div className="col-lg-2 col-lg-offset-2">
							<label >Description</label>

							<textarea
								className="textAreaSize"
								id="description"
								name="Description"
								onChange={ this.handleFields }
								placeholder="Description"
								value={ this.props.props.incidents.description }
							/>

						</div>
					</div>

				</div>
				<div className="clearfix center paddingForm">
					<form onSubmit={ this.handleSubmit } >
						<button className="btn btn-info btn-lg" type="submit">Submit</button>
					</form>
				</div>

			</div>
		)
	}
}

export default reduxForm({
	form: 'Add_incident',// a unique identifier for this form
})(Add_incident)

