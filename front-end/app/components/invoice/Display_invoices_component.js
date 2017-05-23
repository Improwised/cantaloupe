import { BootstrapTable, TableHeaderColumn } from 'react-bootstrap-table';
import React, { Component } from 'react';
import { Link } from 'react-router';
import './styles.css';

function linkToEditInvoice(cell, row) {
	return <Link to={`/invoices/edit/${row.Id}`}>{ cell }</Link>
}

class Display_invoices extends Component {
	componentDidMount() {
		this.props.actions.fetchInvoices(this.props.dispatch); 
	}

	render() {
		var table;
		if (this.props.state.invoices.Invoices && this.props.state.invoices.Invoices.length) {
		table = (	<BootstrapTable
								data={ this.props.state.invoices.Invoices }
								exportCSV={true}
								pagination={true}
								search={true}
								striped={true}
								hover={true}>
								<TableHeaderColumn width="60" isKey={true} dataSort={true} dataField="Id" hidden={true}>#</TableHeaderColumn>
								<TableHeaderColumn dataSort={true} dataField="Invoice_number" dataFormat={linkToEditInvoice} >Invoice Number</TableHeaderColumn>
								<TableHeaderColumn dataSort={true} dataField="Invoice_date">Invoice Date(YYYY-MM-DD)</TableHeaderColumn>
								<TableHeaderColumn dataSort={true} dataField="Invoicer_name">Invoicer</TableHeaderColumn>
							</BootstrapTable>);
		} else {
			table = (<div><div className="panel b block-center text-center"> <h3> You do not have any Data </h3> </div> </div>)
		}
		return (
			<div>
				<div className="paddingLeft"><Link to={`invoices/add/`}><button type="button" className="btn btn-info">Add Invoice</button></Link></div>
				<div>
					{table}
				</div>
			</div>
		)
	}
}

export default Display_invoices
