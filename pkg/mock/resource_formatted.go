package mock

const (
	// VerboseResouce is a long-ass string
	VerboseResouce = `lyra::aws::instance:
  - 'myapp-instance':
      region: 'eu-west-1'
      image_id: 'ami-f90a4880'
      key_name: 'myapp-keypair'
      ensure: 'present'
      additional_info: ''
      block_device_mappings:
      client_token: ''
      cpu_options:
        lyra::aws::cpuoptions:
            core_count: 1
            threads_per_core: 1
      iam_instance_profile:
        lyra::aws::iaminstanceprofile:
          arn: ''
          name: ''
          id: ''
      instance_initiated_shutdown_behavior: ''
      instance_type: 't2.nano'
      ipv6_address_count: 0
      ipv6_addresses:
      kernel_id: ''
      launch_template:
        lyra::aws::launchtemplatespecification:
          launch_template_id: ''
          launch_template_name: ''
          version: ''
      monitoring:
        lyra::aws::monitoring:
          enabled: false
          state: ''
      network_interfaces:
        - lyra::aws::instancenetworkinterface:
            associate_public_ip_address: false
            delete_on_termination: false
            description: ''
            device_index: 0
            subnet_id: 'subnet-07725cdd9bebd4b2a'
            ipv6_address_count: 0
            ipv6_addresses:
            network_interface_id: 'eni-020ba2ff4d5a56f17'
            private_ip_address: '192.168.1.114'
            private_ip_addresses:
              lyra::aws::instanceprivateipaddress:
                primary: true
                private_ip_address: '192.168.1.114'
                association:
                  lyra::aws::instancenetworkinterfaceassociation:
                    ip_owner_id: ''
                    public_dns_name: ''
                    public_ip: ''
                    private_dns_name: 'ip-192-168-1-114.eu-west-1.compute.internal'
            secondary_private_ip_address_count: 0
            association:
              lyra::aws::instancenetworkinterfaceassociation:
                ip_owner_id: 'amazon'
                public_dns_name: 'ec2-34-245-214-144.eu-west-1.compute.amazonaws.com'
                public_ip: '34.245.214.144'
            attachment:
              lyra::aws::instancenetworkinterfaceattachment:
                attach_time: '2018-08-13 14:51:29 +0000 UTC'
                attachment_id: 'eni-attach-027e44030fc0916bc'
                delete_on_termination: true
                device_index: 0
                status: 'attached'
            groups:
              - lyra::aws::groupidentifier:
                  group_id: 'sg-0c5c9c8c1e705ed21'
                  group_name: 'myapp-secgroup'
            mac_address: '06:46:32:fb:3a:a4'
            owner_id: '016439112926'
            private_dns_name: 'ip-192-168-1-114.eu-west-1.compute.internal'
            source_dest_check: true
            status: 'in-use'
            vpc_id: 'vpc-0b006cf2254825f14'
      placement:
        lyra::aws::placement:
          affinity: ''
          availability_zone: 'eu-west-1b'
          group_name: ''
          host_id: ''
          spread_domain: ''
          tenancy: 'default'
      private_ip_address: '192.168.1.114'
      ramdisk_id: ''
      subnet_id: 'subnet-07725cdd9bebd4b2a'
      tag_specifications:
        lyra::aws::tagspecification:
          resource_type: 'instance'
          tags:
            name: 'myapp-instance'
            created_by: 'admin@example.com'
            department: 'engineering'
            lifetime: '1h'
            nyx-logical-id: '1196982c975518b432dbd5fbe644068d1b219b20296ead901bfc6582868879d3'
            project: 'incubator'
            termination_date: '2018-08-13T15:51:30.954179+00:00'
      user_data: ''
      owner_id: '016439112926'
      requester_id: ''
      reservation_id: 'r-08081d11e311e8827'
      ami_launch_index: 0
      architecture: 'x86_64'
      ena_support: true
      hypervisor: 'xen'
      instance_id: 'i-003ae42756a845bac'
      instance_lifecycle: ''
      platform: ''
      private_dns_name: 'ip-192-168-1-114.eu-west-1.compute.internal'
      product_codes:
      public_dns_name: 'ec2-34-245-214-144.eu-west-1.compute.amazonaws.com'
      public_ip_address: '34.245.214.144'
      ram_disk_id: ''
      root_device_name: '/dev/sda1'
      root_device_type: 'ebs'
      security_groups:
        - lyra::aws::groupidentifier:
            group_id: 'sg-0c5c9c8c1e705ed21'
            group_name: 'myapp-secgroup'
      source_dest_check: true
      spot_instance_request_id: ''
      sriov_net_support: ''
      state:
        lyra::aws::instancestate:
          code: 16
          name: 'running'
      state_reason:
        lyra::aws::statereason:
          code: ''
          message: ''
      state_transition_reason: ''
      tags:
        - lyra::aws::tag:
            key: 'department'
            value: 'engineering'
        - lyra::aws::tag:
            key: 'lifetime'
            value: '1h'
        - lyra::aws::tag:
            key: 'Name'
            value: 'myapp-instance'
        - lyra::aws::tag:
            key: 'nyx-logical-id'
            value: '1196982c975518b432dbd5fbe644068d1b219b20296ead901bfc6582868879d3'
        - lyra::aws::tag:
            key: 'termination_date'
            value: '2018-08-13T15:51:30.954179+00:00'
        - lyra::aws::tag:
            key: 'project'
            value: 'incubator'
        - lyra::aws::tag:
            key: 'created_by'
            value: 'admin@example.com'
      virtualization_type: 'hvm'
      vpc_id: 'vpc-0b006cf2254825f14'
`
)
