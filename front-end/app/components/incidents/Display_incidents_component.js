import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import React, {Component} from 'react';
import { Link } from 'react-router';
import Select from 'react-select';
import 'react-select/dist/react-select.css';
import './styles.css';

function incidentInformation(cell, row) {
	return <Link to={`/incidents/${row.Id}`}>{ cell }</Link>
}

function Validator(value) {
	if(!value) {
		return "*required";
	}
	return true;
}

let Components = [];
let ComponentIds = [];
class Display_incidents extends Component {
	constructor(props) {
		super(props);
		this.state = {
			activeAll: 'active'
		};
		this.filterIncidents = this.filterIncidents.bind(this);
		this.editIncident = this.editIncident.bind(this);
	}

	editIncident(row, cellName, cellValue){
		let indexOfComponentId = Components.indexOf(cellValue);
		let ComponentId = ComponentIds[indexOfComponentId]
		this.props.actions.editIncident(row, ComponentId);
	}

	filterIncidents(val) {
		let all = false;
		if(val.value == "all") {
			all = true
		}
		this.setState({activeAll:val.value});
		if(this.state.activeAll != val.value) {
			this.props.actions.fetchIncidents(all);
		}
	}

	componentDidMount() {
		this.props.actions.fetchIncidents(false); 
		this.props.actions.fetchComponents(false);
	}

	render() {
		let options = [
			{ value: 'active', label: 'Active' },
			{ value: 'all', label: 'All' }
		];

		let ComponentId;
		for(let i = 0; i < this.props.state.components.Components.length; i++) {
			ComponentIds[i] = this.props.state.components.Components[i].Id;
			Components[i] = this.props.state.components.Components[i].Name;
		}

		let table;
		if (this.props.state.incidents.Incidents && this.props.state.incidents.Incidents.length) {
		table = (	<BootstrapTable
								data={this.props.state.incidents.Incidents}
								pagination={true}
								search={true}
								striped={true}
								exportCSV={true}
								cellEdit={{
									mode: "dbclick",
									blurToSave: true,
									afterSaveCell: this.editIncident
                }}
								hover={true}>
								<TableHeaderColumn width="50"  dataSort={true} dataField="Id" editable={false} isKey={true} hidden={true}>#</TableHeaderColumn>
								<TableHeaderColumn width="120" dataSort={true} dataField="Component" editable={{type:'select', options:{ values:Components }}}>Component</TableHeaderColumn>
								<TableHeaderColumn width="200" dataSort={true} dataField="Title" dataFormat={incidentInformation} >Title</TableHeaderColumn>
								<TableHeaderColumn width="300" dataSort={true} dataField="Description">Description</TableHeaderColumn>
								<TableHeaderColumn width="160"  dataSort={true} dataField="Warranty_till" editable={false}>Warranty(YYYY-MM-DD)</TableHeaderColumn>
								<TableHeaderColumn width="80"  dataSort={true} dataField="Recorder">Recorder</TableHeaderColumn>
								<TableHeaderColumn width="80"  dataSort={true} dataField="Status" editable={false}>Status</TableHeaderColumn>
							</BootstrapTable>);
		} else {
			table = (<div><div className="panel b block-center text-center"> <h3> You do not have any Data </h3> </div> </div>)
		}

		return (
			<div>
				<div className="clearfix">
					<div className="col-lg-2">
						<Select
							clearable={ false }
							className="activeStyle"
							options={ options }
							onChange={ this.filterIncidents }
							placeholder="Active"
							searchable={ false }
							value={ this.state.activeAll }
						/>
					</div>
					<div className="pull-right">
					 <Link to={`/incidents/add`}><button type="button" className="btn btn-info marginLeftRecord">Record New Incident</button></Link>
					</div>
				</div>
				<div>
					{ table }
				</div>
			</div>
		)
	}
}

export default Display_incidents
