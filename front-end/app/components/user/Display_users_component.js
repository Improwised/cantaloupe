import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import React, {Component} from 'react';
import { Link } from 'react-router'
import axios from 'axios';
import Select from 'react-select';
import 'react-select/dist/react-select.css';
import './styles.css';

var selectRowProp = {
	mode: "checkbox",
	clickToSelect: true,
	bgColor: "rgb(238, 193, 213)"
};

function Validator(value){
	if(!value){
		return '*required'
	}
	return true;
}

function Validator(value){
	if(!value){
		return '*required'
	}
	if(this.type == 'email') {
		var re = /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
		if((re.test(value)) == false) {
			return 'invalid email'
		}
	}
	return true;
}

let MachineNames = [];
let MachineIds = [];
class Display_users extends Component {
	constructor(props) {
		super(props);
		this.state = {
			data :[],
			activeAll: 'active'
		};
		this.filterUsers = this.filterUsers.bind(this);
		this.editUser = this.editUser.bind(this);
		this.deleteUser = this.deleteUser.bind(this);
		this.addUser = this.addUser.bind(this);
	}

	editUser(row, cellName, cellValue) {
		let indexOfMachineId = MachineNames.indexOf(row.Machine_name);
		let MachineId = MachineIds[indexOfMachineId]
		this.props.actions.editUser(row, MachineId);
	}

	addUser(row) {
		let indexOfMachineId = MachineNames.indexOf(row.Machine_name);
		let MachineId = MachineIds[indexOfMachineId]
		this.props.actions.addUser(row, MachineId);
	}

	deleteUser(id) {
		this.props.actions.deleteUser(id);
	}

	filterUsers(val) {
		this.setState({activeAll: val.value});
		var all = false;
		if(val.value == "all") {
			all = true
		}
		if(this.state.activeAll != val.value)
			this.props.actions.fetchUsers(all);
	}

	componentDidMount() {
		this.props.actions.fetchUsers(false);
		this.props.actions.fetchMachines(false);
		$('.modal-content').css("width", "500px");
	}

	render() {

		let width = {
			width: '500px'
		}

		let options = [
			{ value: 'active', label: 'Active' },
			{ value: 'all', label: 'All' }
		];

		let MachineId;
		for(let i = 0; i < this.props.state.machines.Machines.length; i++) {
			MachineIds[i] = this.props.state.machines.Machines[i].Id;
			MachineNames[i] = this.props.state.machines.Machines[i].Name;
		}

		let users = [];
		if(this.props.state.users.AllUsers.length) {
			users = this.props.state.users.AllUsers;
		}

		return (
			<div>
				<div className="clearfix">
					<div className="col-lg-2">

						<Select
							searchable={ false }
							className="activeStyle"
							clearable={ false }
							placeholder="Active"
							value={ this.state.activeAll }
							options={ options }
							onChange={ this.filterUsers }
						/>

					</div>
				</div>
				<div>
					<BootstrapTable
						data={ users }
						pagination={true}
						options={{
							afterDeleteRow :this.deleteUser,
							onAddRow :this.addUser
						}}
						deleteRow={true}
						selectRow={selectRowProp}
						insertRow={true}
						exportCSV={true} // to export data in CSV format ....
						cellEdit={{
							mode: "dbclick",
							blurToSave: true,
							afterSaveCell: this.editUser
						}}
						search={true}
						striped={true}
						hover={true}>
						<TableHeaderColumn width="60"  dataSort={true} dataField="Id" editable={false} isKey={true} autoValue={true} hidden={true} >Id</TableHeaderColumn>
						<TableHeaderColumn width="260" dataSort={true} dataField="Name" editable={{ validator:Validator }} >User</TableHeaderColumn>
						<TableHeaderColumn width="350" dataSort={true} dataField="Company_email" editable={{ type: 'email', validator:Validator }}>E-mail</TableHeaderColumn>
						<TableHeaderColumn width="150" dataSort={true} dataField="Machine_name" editable={{type:'select', options:{ values:MachineNames }}}>Machine</TableHeaderColumn>
					</BootstrapTable>
				</div>
			</div>
		)
	}
}

export default Display_users
