This contains resources for testing that the MCP server tools integrate properly with an LLM

~~~bash
cd integration_testing
podman-compose up
~~~

After a few seconds, an output similar to this should appear :

~~~
[app]    | ####################################
[app]    | [{'role': 'system', 'content': "You're a helpful assistant for a database system administrator"}, {'role': 'user', 'content': 'Can you tell me how many roles there are in the cluster ?'}, {'role': 'assistant', 'tool_calls': [{'id': 'call_aalr01n8', 'function': Function(arguments='{}', name='list_all_roles'), 'type': 'function'}]}, {'role': 'tool', 'tool_call_id': 'call_aalr01n8', 'name': 'list_all_roles', 'content': '[pg_database_owner pg_read_all_data pg_write_all_data pg_monitor pg_read_all_settings pg_read_all_stats pg_stat_scan_tables pg_read_server_files pg_write_server_files pg_execute_server_program pg_signal_backend pg_checkpoint pg_maintain pg_use_reserved_connections pg_create_subscription postgres]'}, {'role': 'assistant', 'content': 'There are 15 roles in the cluster.'}]
~~~

