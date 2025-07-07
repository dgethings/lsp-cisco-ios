-- Neovim snippets for Cisco IOS configurations

return {
  -- Interface Configuration
  interface = {
    prefix = "interface",
    body = {
      "interface ${1}",
      "  description ${2}",
      "  ip address ${3} ${4}",
      "  no shutdown",
    },
    description = "Cisco IOS Interface Configuration",
  },
  -- BGP Router Configuration
  router_bgp = {
    prefix = "router bgp",
    body = {
      "router bgp ${1}",
      "  bgp router-id ${2}",
      "  neighbor ${3} remote-as ${4}",
      "  network ${5} mask ${6}",
    },
    description = "Cisco IOS BGP Router Configuration",
  },
  -- Static Route
  ip_route = {
    prefix = "ip route",
    body = {
      "ip route ${1} ${2} ${3}",
    },
    description = "Cisco IOS Static Route",
  },
}
