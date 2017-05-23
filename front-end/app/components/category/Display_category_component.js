import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import React, {Component} from 'react';
import { Link } from 'react-router'
import axios from 'axios';
import Select from 'react-select';
import 'react-select/dist/react-select.css';
import './styles.css';

function Validator(value){
	if(!value){
		return "*required";
	}
	return true;
}

class Display_category extends Component {
	constructor(props) {
		super(props);
		this.state = {
			data :[],
			activeAll: 'active'
		};
		this.filterCategory = this.filterCategory.bind(this);
		this.editCategory = this.editCategory.bind(this);
		this.deleteCategory = this.deleteCategory.bind(this);
		this.addCategory = this.addCategory.bind(this);
	}

	editCategory(row, cellName, cellValue) {
		this.props.actions.editCategory(row);
	}

	addCategory(row) {
		this.props.actions.addCategory(row);
	}

	deleteCategory(id) {
		this.props.actions.deleteCategory(id);
	}

	filterCategory(val) {
		this.setState({activeAll: val.value});
		var all = false;
		if(val.value == "all") {
			all = true
		}
		if(this.state.activeAll != val.value) {
			this.props.actions.fetchCategories(all);
		}
	}

	componentDidMount() {
		this.props.actions.fetchCategories(false);
		$('.modal-content').css("width", "500px");
	}

	render() {
		let categories = [];
		if(this.props.state.category.AllCategories.length) {
			categories = this.props.state.category.AllCategories;
		}

		let selectRowProp = {
			mode: "checkbox",
			clickToSelect: true,
			bgColor: "rgb(238, 193, 213)"
		};

		let options = [
			{	value: 'active', label: 'Active' },
			{	value: 'all', label: 'All' }
		];

		return (
			<div>
				<div className="clearfix">
					<div className="col-lg-2">
						<Select
							className="activeStyle"
							clearable={ false }
							options={ options }
							onChange={ this.filterCategory }
							placeholder="Active"
							searchable={ false }
							value={ this.state.activeAll }
						/>
					</div>
				</div>
				<div>

					<BootstrapTable
						data={ categories }
						pagination={true}
						options={{
							afterDeleteRow :this.deleteCategory,
							onAddRow :this.addCategory
						}}
						deleteRow={true}
						selectRow={selectRowProp}
						insertRow={true}
						exportCSV={true}
						cellEdit={{
							mode: "dbclick",
							blurToSave: true,
							afterSaveCell: this.editCategory
						}}
						search={true}
						striped={true}
						hover={true} >
						<TableHeaderColumn dataField="Id" editable={false} isKey={true} autoValue={true} hidden={true} >Id</TableHeaderColumn>
						<TableHeaderColumn width="260" dataSort={true} dataField="Category" editable={{ validator:Validator }} >Category</TableHeaderColumn>
						<TableHeaderColumn width="350" dataSort={true} dataField="Description" editable={{ type:'textarea' }} >Description</TableHeaderColumn>
					</BootstrapTable>

				</div>
			</div>
		)
	}
}

export default Display_category
