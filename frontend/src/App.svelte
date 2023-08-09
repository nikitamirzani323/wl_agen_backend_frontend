<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navbar from "./components/Navbar.svelte";
	import Dashboard from "./pages/dashboard/Dashboard.svelte";
	
	import Admin from "./pages/admin/Admin.svelte";
	import Adminrule from "./pages/adminrule/Adminrule.svelte";
	import Provider from "./pages/provider/Provider.svelte";
	import Currency from "./pages/currency/Currency.svelte";
	import Catebank from "./pages/catebank/Catebank.svelte";
	import Categame from "./pages/categame/Categame.svelte";
	import Master from "./pages/master/Master.svelte";
	import Login from "./pages/Login.svelte";
	import NotFound from "./pages/NotFound.svelte";
	export let table_header_font;
	export let table_body_font;
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
				component: Dashboard,
			}),
			"/categame": wrap({
				component: Categame,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/provider": wrap({
				component: Provider,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/master": wrap({
				component: Master,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/catebank": wrap({
				component: Catebank,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/currency": wrap({
				component: Currency,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			
			"/admin": wrap({
				component: Admin,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/adminrule": wrap({
				component: Adminrule,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navbar />
{/if}
<Router {routes} />
