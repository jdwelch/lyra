package mock

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
)

// DynamicFeedbackOnProgress is a sketch of a
func DynamicFeedbackOnProgress() {
	ui.DescribeStepWithField("Applying manifest:", "/home/ubuntu/nyx-examples/aws_ec2_puppet/ec2_instance_in_vpc_create.pp")

	ui.ProgressBar("Applying resource: Lyra::Aws::Importkeypair['myapp-keypair']", 2000, false)
	ui.ResourceSet("Lyra::Aws::Importkeypair('title' => 'myapp-keypair', 'region' => 'eu-west-1', 'public_key_material' => 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCX363gh/q6DGSL963/LlYcILkYKtEjrq5Ze4gr1BJdY0pqLMIKFt/VMJ5UTyx85N4Chjb/jEQhZzlWGC1SMsXOQ+EnY72fYrpOV0wZ4VraxZAz3WASikEglHJYALTQtsL8RGPxlBhIv0HpgevBkDlHvR+QGFaEQCaUhXCWDtLWYw== nyx-test-keypair-nopassword', 'ensure' => 'present', 'key_fingerprint' => 'c5:5d:ee:ed:13:c4:da:46:d9:88:a4:8a:34:e9:1f:52', 'key_name' => 'myapp-keypair')")
	ui.Notice("notice: Imported KeyPair: myapp-keypair, FingerPrint: c5:5d:ee:ed:13:c4:da:46:d9:88:a4:8a:34:e9:1f:52\n")

	ui.ProgressBar("Applying resource: Lyra::Aws::Vpc[myapp-vpc]", 2000, false)
	ui.ResourceSet("Lyra::Aws::Vpc('title' => 'myapp-vpc', 'region' => 'eu-west-1', 'cidr_block' => '192.168.0.0/16', 'ensure' => 'present', 'enable_dns_hostnames' => true, 'enable_dns_support' => true, 'tags' => {'Name' => 'myapp-vpc', 'created_by' => 'admin@example.com', 'department' => 'engineering', 'lifetime' => '1h', 'nyx-logical-id' => '2f3c4722dc9ac37e6aa902088e564862f3c404ad8edee0814d116694d4300b33', 'project' => 'incubator'}, 'vpc_id' => 'vpc-0936d7925b4cb6c75', 'is_default' => false, 'state' => 'available', 'dhcp_options_id' => 'dopt-58a6ba3a')\n")
	ui.Notice("Created VPC: vpc-0936d7925b4cb6c75")

	ui.ProgressBar("Applying resource: Lyra::Aws::Securitygroup[myapp-secgroup]", 2000, false)
	ui.ResourceSet("Lyra::Aws::Securitygroup('title' => 'myapp-secgroup', 'region' => 'eu-west-1', 'description' => 'myapp-secgroup', 'ensure' => 'present', 'group_name' => 'myapp-secgroup', 'vpc_id' => 'vpc-0936d7925b4cb6c75', 'group_id' => 'sg-0817f65146f7abacf', 'ip_permissions' => [Lyra::Aws::Ippermission('from_port' => 0, 'ip_protocol' => '-1', 'to_port' => 0, 'ip_ranges' => [Lyra::Aws::Iprange('cidr_ip' => '0.0.0.0/0', 'description' => 'any source address')])], 'ip_permissions_egress'	=> [Lyra::Aws::Ippermission('from_port' => 0, 'ip_protocol' => '-1', 'to_port' => 0, 'ip_ranges' => [Lyra::Aws::Iprange('cidr_ip' => '0.0.0.0/0', 'description' => '')])], 'owner_id' => '016439112926', 'tags' => {'Name' => 'myapp-secgroup', 'created_by' => 'admin@example.com', 'department' => 'engineering', 'lifetime' => '1h', 'nyx-logical-id' => '603883a4f67e9dcaab5e72c2d20a15378bb21daa69a1f2ab20fc36ed519d73c8', 'project' => 'incubator'})\n")

	ui.Notice("notice: Created SecurityGroup: sg-0817f65146f7abacf")

	ui.ProgressBar("Applying resource: Lyra::Aws::Subnet[myapp-subnet]", 2000, false)
	ui.ResourceSet("Lyra::Aws::Subnet('title' => 'myapp-subnet', 'region' => 'eu-west-1', 'vpc_id' => 'vpc-0936d7925b4cb6c75','cidr_block' => '192.168.1.0/24', 'ensure' => 'present', 'availability_zone' => 'eu-west-1a', 'ipv6_cidr_block' => '', 'tags' => {'Name' => 'myapp-subnet', 'created_by' => 'admin@example.com', 'department' => 'engineering', 'lifetime' => '1h', 'nyx-logical-id' => '90ab2286ae213203e94e54a98bdeac96ae0b31695b6bd50b01cb6ab95851b4fa', 'project' => 'incubator'}, 'assign_ipv6_address_on_creation' => false, 'map_public_ip_on_launch' => true, 'available_ip_address_count' => 251, 'default_for_az' => false, 'state' => 'available', 'subnet_id' => 'subnet-0c57a82edffceb552')\n")
	ui.Notice("notice: Created Subnet: subnet-0c57a82edffceb552")

	ui.ProgressBar("Applying resource: Lyra::Aws::Internetgateway[myapp-gateway]", 2000, false)
	ui.ResourceSet("Lyra::Aws::Internetgateway('title' => 'myapp-gateway', 'region' => 'eu-west-1', 'ensure' => 'present', 'internet_gateway_id' => 'igw-0c0bdeeb34cfc2534', 'tags' => {'Name' => 'myapp-gateway', 'created_by' => 'admin@example.com', 'department' => 'engineering', 'lifetime' => '1h', 'nyx-logical-id' => 'a3bfa9bb2728b4ccb64598d81081718b42dad58e1ceb6b4811c3dfd95430047a', 'project' => 'incubator'}, 'attachments' => [])\n")

	ui.ProgressBar("Applying resource: Lyra::Aws::Attachinternetgateway[myapp-gateway]", 2000, false)
	ui.ResourceSet("Lyra::Aws::Attachinternetgateway('title' => 'myapp-gateway', 'region' => 'eu-west-1', 'internet_gateway_id'	=> 'igw-0c0bdeeb34cfc2534', 'vpc_id' => 'vpc-0936d7925b4cb6c75', 'ensure' => 'present')\n")
	ui.Notice("notice: Attached Internet Gateway: igw-0c0bdeeb34cfc2534 to VPC vpc-0936d7925b4cb6c75")

	ui.ProgressBar("Applying resource: Lyra::Aws::Routetable[myapp-routes]", 2000, false)
	ui.ResourceSet("Lyra::Aws::Routetable('title' => 'myapp-routes', 'region' => 'eu-west-1', 'vpc_id' => 'vpc-0936d7925b4cb6c75', 'ensure' => 'present', 'route_table_id' => 'rtb-0f2eafd5eb42f1447', 'subnet_id' => '', 'routes' => [Lyra::Aws::Route('destination_cidr_block' => '192.168.0.0/16', 'destination_ipv6_cidr_block' => '', 'destination_prefix_list_id' => '', 'egress_only_internet_gateway_id' => '', 'gateway_id' => 'local', 'instance_id' => '', 'instance_owner_id' => '', 'nat_gateway_id' => '', 'network_interface_id' => '', 'origin' => 'CreateRouteTable', 'state' => 'active', 'vpc_peering_connection_id' => ''), Lyra::Aws::Route('destination_cidr_block' => '0.0.0.0/0', 'destination_ipv6_cidr_block' => '', 'destination_prefix_list_id' => '', 'egress_only_internet_gateway_id' => '', 'gateway_id' => 'igw-0c0bdeeb34cfc2534', 'instance_id' => '', 'instance_owner_id' => '', 'nat_gateway_id' => '', 'network_interface_id' => '', 'origin' => 'CreateRoute', 'state' => 'active', 'vpc_peering_connection_id' => '')], 'associations' => [Lyra::Aws::Routetableassociation('main' => false, 'route_table_association_id' => 'rtbassoc-0dedfabfee72055a4', 'route_table_id' => 'rtb-0f2eafd5eb42f1447', 'subnet_id' => 'subnet-0c57a82edffceb552')], 'tags' => {'Name' => 'myapp-routes', 'created_by' => 'admin@example.com', 'department' => 'engineering', 'lifetime' => '1h', 'nyx-logical-id' => '71f6f40edaf695bd7a0ff56867ae3a9f1ac80b0b910628d0ab7b6484de4abe1a', 'project' => 'incubator'})\n")

	ui.ProgressBar("Applying resource: Lyra::Aws::Instance[myapp-instance]", 4000, false)
	ui.ResourceSet("Lyra::Aws::Instance('title' => 'myapp-instance', 'region' => 'eu-west-1', 'image_id' => 'ami-f90a4880', 'key_name' => 'myapp-keypair', 'ensure' => 'present', 'additional_info' => '', 'block_device_mappings' => [], 'client_token' =>	'', 'cpu_options' => Lyra::Aws::Cpuoptions('core_count' => 1, 'threads_per_core' => 1), 'iam_instance_profile' => Lyra::Aws::Iaminstanceprofile('arn' => '', 'name' => '', 'id' => ''), 'instance_initiated_shutdown_behavior' => '', 'instance_type' => 't2.nano', 'ipv6_address_count' => 0, 'ipv6_addresses' => [], 'kernel_id' => '', 'launch_template' => Lyra::Aws::Launchtemplatespecification('launch_template_id' => '', 'launch_template_name' => '', 'version' => ''), 'monitoring' => Lyra::Aws::Monitoring('enabled' => false, 'state' => ''), 'network_interfaces' => [Lyra::Aws::Instancenetworkinterface('associate_public_ip_address' => false, 'delete_on_termination' => false, 'description' => '', 'device_index' => 0, 'subnet_id' => 'subnet-0c57a82edffceb552', 'ipv6_address_count' => 0, 'ipv6_addresses' => [], 'network_interface_id' => 'eni-005baadf7cd27559d', 'private_ip_address' => '192.168.1.122', 'private_ip_addresses' => [Lyra::Aws::Instanceprivateipaddress('primary' => true, 'private_ip_address' => '192.168.1.122', 'association' => Lyra::Aws::Instancenetworkinterfaceassociation('ip_owner_id' => '', 'public_dns_name' => '', 'public_ip' => ''), 'private_dns_name' => 'ip-192-168-1-122.eu-west-1.compute.internal')], 'secondary_private_ip_address_count' => 0, 'association' => Lyra::Aws::Instancenetworkinterfaceassociation('ip_owner_id' =>'amazon', 'public_dns_name' => 'ec2-34-247-54-147.eu-west-1.compute.amazonaws.com', 'public_ip' => '34.247.54.147'), 'attachment' => Lyra::Aws::Instancenetworkinterfaceattachment('attach_time' => '2018-09-03 14:46:41 +0000 UTC', 'attachment_id' => 'eni-attach-09cc3c1dcb9917671', 'delete_on_termination' => true, 'device_index' => 0, 'status' => 'attached'), 'groups' => [Lyra::Aws::Groupidentifier('group_id' => 'sg-0817f65146f7abacf', 'group_name' => 'myapp-secgroup')], 'mac_address' => '02:bb:c8:b3:da:b8', 'owner_id' => '016439112926', 'private_dns_name' => 'ip-192-168-1-122.eu-west-1.compute.internal', 'source_dest_check' => true, 'status' => 'in-use', 'vpc_id' => 'vpc-0936d7925b4cb6c75')], 'placement' => Lyra::Aws::Placement('affinity' => '', 'availability_zone' => 'eu-west-1a', 'group_name' => '', 'host_id' => '', 'spread_domain' => '', 'tenancy' => 'default'), 'private_ip_address' => '192.168.1.122', 'ramdisk_id' => '', 'subnet_id' => 'subnet-0c57a82edffceb552', 'tag_specifications' => [Lyra::Aws::Tagspecification('resource_type' => 'instance', 'tags' => {'Name' => 'myapp-instance', 'created_by' => 'admin@example.com', 'department' => 'engineering', 'lifetime' => '1h', 'nyx-logical-id' => '1196982c975518b432dbd5fbe644068d1b219b20296ead901bfc6582868879d3', 'project' => 'incubator', 'termination_date' => '2018-09-03T15:46:42.884339+00:00'})], 'user_data' => '', 'owner_id' => '016439112926', 'requester_id' => '', 'reservation_id' => 'r-02735c7afd3f9cfde', 'ami_launch_index' => 0, 'architecture' => 'x86_64', 'ena_support' => true, 'hypervisor' => 'xen', 'instance_id' => 'i-0ceef56fbb65e5971', 'instance_lifecycle' => '', 'platform' => '', 'private_dns_name' => 'ip-192-168-1-122.eu-west-1.compute.internal', 'product_codes' => [], 'public_dns_name' => 'ec2-34-247-54-147.eu-west-1.compute.amazonaws.com', 'public_ip_address' => '34.247.54.147', 'ram_disk_id' => '', 'root_device_name' => '/dev/sda1', 'root_device_type' => 'ebs', 'security_groups' => [Lyra::Aws::Groupidentifier('group_id' => 'sg-0817f65146f7abacf', 'group_name' => 'myapp-secgroup')], 'source_dest_check' => true, 'spot_instance_request_id' => '', 'sriov_net_support' => '', 'state' => Lyra::Aws::Instancestate('code' => 16, 'name' => 'running'), 'state_reason' => Lyra::Aws::Statereason('code' => '', 'message' => ''), 'state_transition_reason' => '', 'tags'=> [Lyra::Aws::Tag('key' => 'created_by', 'value' => 'admin@example.com'), Lyra::Aws::Tag('key' => 'department', 'value' => 'engineering'), Lyra::Aws::Tag('key' => 'Name', 'value' => 'myapp-instance'), Lyra::Aws::Tag('key' => 'nyx-logical-id', 'value' => '1196982c975518b432dbd5fbe644068d1b219b20296ead901bfc6582868879d3'), Lyra::Aws::Tag('key' => 'termination_date', 'value' => '2018-09-03T15:46:42.884339+00:00'), Lyra::Aws::Tag('key' => 'lifetime', 'value' => '1h'), Lyra::Aws::Tag('key' => 'project', 'value' => 'incubator')], 'virtualization_type' => 'hvm', 'vpc_id' => 'vpc-0936d7925b4cb6c75')\n")

	ui.Notice("notice: Created EC2 Instance, ID: i-0ceef56fbb65e5971, PrivateIP: 192.168.1.122, PublicIP: 34.247.54.147")

	ui.DescribeStepWithField("Success!", "apply finished")
}

// TableProgress is a sketch of a potential solution for
func TableProgress() {
	fmt.Print(`resource                             title                     status
--------------------------------------------------------------------------------
lyra::aws::importkeypair          'myapp-keypair-long-title'  done
lyra::aws::vpc                    'myapp-vpc'                 done
lyra::aws::securitygroup          'myapp-secgroup'            working…
lyra::aws::subnet                 'myapp-subnet'              waiting…
lyra::aws::internetgateway        'myapp-gateway'             waiting…
lyra::aws::attachinternetgateway  'myapp-gateway'             waiting…
lyra::aws::routetable             'myapp-routes'              waiting…
lyra::aws::instance               'myapp-instance'            waiting…
lyra::ssh::exec                   'myapp-instance'            waiting…
--------------------------------------------------------------------------------
`)
}
