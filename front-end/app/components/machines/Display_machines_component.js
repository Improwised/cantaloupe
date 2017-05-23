import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import React, {Component} from 'react';
import { Link } from 'react-router';
import Select from 'react-select';
import 'react-select/dist/react-select.css';
import './styles.css';

function Validator(value){
	if(!value){
		return "*required"
	}
	return true;
}

function machineInformation(cell, row){
	 return <Link to={`/machines/${row.Id}/components`}>{ cell }</Link>
}

let userNames = [];
let userIds = [];
class Display_machines extends Component {
	constructor(props) {
		super(props);
		this.state = {
			data: [],
			activeAll: 'active'
		};
		this.filterMachines = this.filterMachines.bind(this);
		this.editMachine = this.editMachine.bind(this);
		this.deleteMachine = this.deleteMachine.bind(this);
		this.addMachine = this.addMachine.bind(this);
	}

	addMachine(row) {
		this.props.actions.addMachine(row);
	}

	editMachine(row, cellName, cellValue) {
		this.props.actions.editMachine(row);
	}

	deleteMachine(id) {
		this.props.actions.deleteMachine(id);
	}

	filterMachines(val) {
		let all = false;
		if(val.value == "all") {
			all = true
		}
		this.setState({activeAll:val.value});
		if(this.state.activeAll != val.value) {
			this.props.actions.fetchMachines(all);
		}
	}

	componentDidMount() {
		this.props.actions.fetchMachines(false, this.props.dispatch);
		this.props.actions.fetchUsers(false, this.props.dispatch);
		$('.modal-content').css("width", "400px");
	}

	render() {
		let machines = [];
		if(this.props.state.machines.Machines.length) {
			machines = this.props.state.machines.Machines;
		}

		let options = [
			{ value: 'active', label: 'Active' },
			{ value: 'all', label: 'All' }
		];

	//===================== style ================================================
		let selectRowProp = {
			mode: "checkbox",
			clickToSelect: true,
			bgColor: "rgb(238, 193, 213)"
		};
	// ===========================================================================

		let table;
		table = (	<BootstrapTable
								data={ machines }
								pagination={true}
								options={{
									afterDeleteRow: this.deleteMachine,
									onAddRow: this.addMachine
								}}
								exportCSV={true}
								deleteRow={true}
								selectRow={selectRowProp}
								insertRow={true}
								cellEdit={{
									mode: "dbclick",
									blurToSave: true,
									afterSaveCell: this.editMachine
								}}
								search={true}
								striped={true}
								hover={true}>
								<TableHeaderColumn width="60"  dataSort={true} dataField="Id" editable={false} isKey={true}  autoValue={true} hidden={true} >Id</TableHeaderColumn>
								<TableHeaderColumn width="260" dataSort={true} dataField="Name" dataFormat={machineInformation} editable={{ validator:Validator }}>Machine</TableHeaderColumn>
								<TableHeaderColumn width="350" dataSort={true} dataField="User" editable={false} hiddenOnInsert={true} autoValue={false}>Current User</TableHeaderColumn>
							</BootstrapTable> );

		return (
			<div>
				<div className="clearfix">
					<div className="col-lg-2">

						<Select
							searchable={ false }
							clearable={ false }
							placeholder="Active"
							value={ this.state.activeAll }
							options={ options }
							className="activeStyle"
							onChange={ this.filterMachines }
						/>

					</div>
				</div>
				<div>
					{ table }
				</div>
			</div>
		)
	}
}

export default Display_machines
