"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from petstore_api.api_client import ApiClient, Endpoint as _Endpoint
from petstore_api.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from petstore_api.model.user import User


class UserApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.create_user_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [],
                'endpoint_path': '/user',
                'operation_id': 'create_user',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'body',
                ],
                'required': [
                    'body',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'body':
                        (User,),
                },
                'attribute_map': {
                },
                'location_map': {
                    'body': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [],
            },
            api_client=api_client
        )
        self.create_users_with_array_input_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [],
                'endpoint_path': '/user/createWithArray',
                'operation_id': 'create_users_with_array_input',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'body',
                ],
                'required': [
                    'body',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'body':
                        ([User],),
                },
                'attribute_map': {
                },
                'location_map': {
                    'body': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [],
            },
            api_client=api_client
        )
        self.create_users_with_list_input_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [],
                'endpoint_path': '/user/createWithList',
                'operation_id': 'create_users_with_list_input',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'body',
                ],
                'required': [
                    'body',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'body':
                        ([User],),
                },
                'attribute_map': {
                },
                'location_map': {
                    'body': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [],
            },
            api_client=api_client
        )
        self.delete_user_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [],
                'endpoint_path': '/user/{username}',
                'operation_id': 'delete_user',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'username',
                ],
                'required': [
                    'username',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'username':
                        (str,),
                },
                'attribute_map': {
                    'username': 'username',
                },
                'location_map': {
                    'username': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [],
            },
            api_client=api_client
        )
        self.get_user_by_name_endpoint = _Endpoint(
            settings={
                'response_type': (User,),
                'auth': [],
                'endpoint_path': '/user/{username}',
                'operation_id': 'get_user_by_name',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'username',
                ],
                'required': [
                    'username',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'username':
                        (str,),
                },
                'attribute_map': {
                    'username': 'username',
                },
                'location_map': {
                    'username': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/xml',
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.login_user_endpoint = _Endpoint(
            settings={
                'response_type': (str,),
                'auth': [],
                'endpoint_path': '/user/login',
                'operation_id': 'login_user',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'username',
                    'password',
                ],
                'required': [
                    'username',
                    'password',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'username':
                        (str,),
                    'password':
                        (str,),
                },
                'attribute_map': {
                    'username': 'username',
                    'password': 'password',
                },
                'location_map': {
                    'username': 'query',
                    'password': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/xml',
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.logout_user_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [],
                'endpoint_path': '/user/logout',
                'operation_id': 'logout_user',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                },
                'attribute_map': {
                },
                'location_map': {
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [],
            },
            api_client=api_client
        )
        self.update_user_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [],
                'endpoint_path': '/user/{username}',
                'operation_id': 'update_user',
                'http_method': 'PUT',
                'servers': None,
            },
            params_map={
                'all': [
                    'username',
                    'body',
                ],
                'required': [
                    'username',
                    'body',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'username':
                        (str,),
                    'body':
                        (User,),
                },
                'attribute_map': {
                    'username': 'username',
                },
                'location_map': {
                    'username': 'path',
                    'body': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [],
            },
            api_client=api_client
        )

    def create_user(
        self,
        body,
        **kwargs
    ):
        """Create user  # noqa: E501

        This can only be done by the logged in user.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.create_user(body, async_req=True)
        >>> result = thread.get()

        Args:
            body (User): Created user object

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['body'] = \
            body
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.create_user_endpoint.call_with_http_info(**kwargs)

    def create_users_with_array_input(
        self,
        body,
        **kwargs
    ):
        """Creates list of users with given input array  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.create_users_with_array_input(body, async_req=True)
        >>> result = thread.get()

        Args:
            body ([User]): List of user object

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['body'] = \
            body
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.create_users_with_array_input_endpoint.call_with_http_info(**kwargs)

    def create_users_with_list_input(
        self,
        body,
        **kwargs
    ):
        """Creates list of users with given input array  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.create_users_with_list_input(body, async_req=True)
        >>> result = thread.get()

        Args:
            body ([User]): List of user object

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['body'] = \
            body
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.create_users_with_list_input_endpoint.call_with_http_info(**kwargs)

    def delete_user(
        self,
        username,
        **kwargs
    ):
        """Delete user  # noqa: E501

        This can only be done by the logged in user.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.delete_user(username, async_req=True)
        >>> result = thread.get()

        Args:
            username (str): The name that needs to be deleted

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['username'] = \
            username
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.delete_user_endpoint.call_with_http_info(**kwargs)

    def get_user_by_name(
        self,
        username,
        **kwargs
    ):
        """Get user by user name  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_user_by_name(username, async_req=True)
        >>> result = thread.get()

        Args:
            username (str): The name that needs to be fetched. Use user1 for testing.

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            User
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['username'] = \
            username
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.get_user_by_name_endpoint.call_with_http_info(**kwargs)

    def login_user(
        self,
        username,
        password,
        **kwargs
    ):
        """Logs user into the system  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.login_user(username, password, async_req=True)
        >>> result = thread.get()

        Args:
            username (str): The user name for login
            password (str): The password for login in clear text

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            str
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['username'] = \
            username
        kwargs['password'] = \
            password
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.login_user_endpoint.call_with_http_info(**kwargs)

    def logout_user(
        self,
        **kwargs
    ):
        """Logs out current logged in user session  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.logout_user(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.logout_user_endpoint.call_with_http_info(**kwargs)

    def update_user(
        self,
        username,
        body,
        **kwargs
    ):
        """Updated user  # noqa: E501

        This can only be done by the logged in user.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.update_user(username, body, async_req=True)
        >>> result = thread.get()

        Args:
            username (str): name that need to be deleted
            body (User): Updated user object

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['username'] = \
            username
        kwargs['body'] = \
            body
        headers = kwargs.get('headers', {})
        if headers:
            accept = headers.get('accept') or headers.get('Accept')
            if accept and accept in self.gene_download_summary_by_accession_endpoint.headers_map:
                self.gene_download_summary_by_accession_endpoint.headers_map['accept'] = accept

            for key in headers.keys():
                self.gene_download_summary_by_accession_endpoint.headers_map[key] = headers[key]

        return self.update_user_endpoint.call_with_http_info(**kwargs)

