import React, {Component} from 'react'
import { textarea, Field, FieldArray, reduxForm } from 'redux-form'
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import {renderField, rendreTextArea, renderComponents, rendreLabel} from '../../fields/fields'
import { Link } from 'react-router'
import './styles.css';

class Machine_information extends Component {
	constructor(props) {
		super(props);
		this.state = {
			data: [],
			component_information: [],
			incident_information: [],
			past_uses_history: []
		};
	}

	componentDidMount() {
		let { machineId } = this.props.params;
		this.props.actions.fetchMachineInformation(machineId);
	}

	render() {

		//====================== style =============================================

		let borderStyle = {
			border: 'solid',
			borderWidth: '2px',
			padding: '15px'
		};

		//==========================================================================

		//null data will be pass to datatable if database contains no data ....
		let MachineName = '';
		let UsingSince = '';
		let CurrentUser = '';
		let PastUses = [];
		let Components = [];
		let Incidents = [];

		if(this.props.props.machines.MachineInfo.Machine) {
			MachineName = this.props.props.machines.MachineInfo.Machine;
			CurrentUser = this.props.props.machines.MachineInfo.Name;
			UsingSince = this.props.props.machines.MachineInfo.UsingSince;
			PastUses = this.props.props.machines.MachineInfo.PastUses;
			Components = this.props.props.machines.MachineInfo.Components;
			Incidents = this.props.props.machines.MachineInfo.Incidents;
		}
		const options = {
    	sizePerPage: 5  // Showing 5 for the size per page as default
  	};
		const { handleSubmit, pristine, reset, submitting } = this.props
		return (

			<div >
				<h1>{ MachineName }</h1>
				<div style={borderStyle}>
					<div>
						<div className="col-lg-offset-1">
							<Field component={rendreLabel} label="Current User" name="user" /> <Link to={`/users`}> { CurrentUser } </Link>
						</div>
						<div className="col-lg-offset-1">
							<Field component={rendreLabel} label="Using Since" name="usingScince" /> { UsingSince }
						</div>
					</div>
				</div>
				<div>
					<div className="center">
						<h4><b> Component(s) </b></h4>
					</div>
					<div>
						<BootstrapTable
							data={ Components }
							pagination={true}
							options={ options }
							search={true}
							striped={true}
							hover={true}
							paginationSize="5" >
							<TableHeaderColumn width="20"  dataSort={true} dataField="Id" isKey={true}>#</TableHeaderColumn>
							<TableHeaderColumn width="50" dataSort={true} dataField="SerialNo">Serial</TableHeaderColumn>
							<TableHeaderColumn width="100" dataSort={true} dataField="Name">Name</TableHeaderColumn>
							<TableHeaderColumn width="180"  dataSort={true} dataField="Description" >Description</TableHeaderColumn>
							<TableHeaderColumn width="60"  dataSort={true} dataField="Warranty_till" >Warranty Till</TableHeaderColumn>
							<TableHeaderColumn width="60"  dataSort={true} dataField="AddOn" >Add On</TableHeaderColumn>
						</BootstrapTable>
					</div>

					<div className="center">
						<h4><b> Incident(s) </b></h4>
					</div>
					<div>
						<BootstrapTable
							data={ Incidents }
							pagination={true}
							options={ options }
							search={true}
							striped={true}
							hover={true}>
							<TableHeaderColumn width="50"  dataSort={true} dataField="Id" isKey={true}>#</TableHeaderColumn>
							<TableHeaderColumn width="180" dataSort={true} dataField="Title">Title</TableHeaderColumn>
							<TableHeaderColumn width="200" dataSort={true} dataField="Description">Description</TableHeaderColumn>
							<TableHeaderColumn width="80" dataSort={true} dataField="Recorder">Recorder</TableHeaderColumn>
							<TableHeaderColumn width="50"  dataSort={true} dataField="Status" >Status</TableHeaderColumn>
						</BootstrapTable>
					</div>
					<div className="center">
						<h4><b> Past Uses History </b></h4>
					</div>
					<div>
						<BootstrapTable
							data={ PastUses }
							options={ options }
							pagination={true}
							search={true}
							striped={true}
							hover={true}>
							<TableHeaderColumn  dataSort={true} dataField="BeginDate" isKey={true}>Begin</TableHeaderColumn>
							<TableHeaderColumn  dataSort={true} dataField="EndDate">End</TableHeaderColumn>
							<TableHeaderColumn  dataSort={true} dataField="User">User</TableHeaderColumn>
						</BootstrapTable>
					</div>
				</div>
			</div>
		)
	}
}

export default reduxForm({
	form: 'component_Info',// a unique identifier for this form
})(Machine_information)

