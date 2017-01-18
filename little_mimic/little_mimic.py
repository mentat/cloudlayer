from datetime import datetime
from datetime import timedelta

from bottle import request
from bottle import response
from bottle import route
from bottle import run


HOST = "localhost"
PORT = "8080"
COMPUTE_URL = "http://{host}:{port}/nova/v2.1/{project_id}"
NEUTRON_URL = "http://{host}:{port}/neutron/v2.0"


@route('/v3')
def v3():
    resp = {"version":
            {"status": "stable",
             "updated": "2016-04-04T00:00:00Z",
             "media-types": [
                 {"base": "application/json",
                  "type": "application/vnd.openstack.identity-v3+json"}
             ],
             "id": "v3.6",
             "links": [
                 {"href": "http://{host}:{port}/v3/".format(host=HOST, port=PORT),
                  "rel": "self"}
             ]}}
    return resp


@route('/v3/auth/tokens', method='POST')
def v3_auth_tokens():
    project_id = request.json['auth']['scope']['project']['id']
    user = request.json['auth']['identity']['password']['user']['name']
    now = datetime.utcnow()
    expires = now + timedelta(hours=2)
    expires_at = "{0}Z".format(expires.isoformat())
    issued_at = "{0}Z".format(now.isoformat())

    response.status = 201
    resp = {"token":
            {"methods": ["password"],
             "roles": [{"id": "71f691fb48aa4eebb27a9ca72c695f6f", "name": "admin"}],
             "expires_at": expires_at,  # "2017-01-17T05:20:17.000000Z",
             "project": {"domain": {"id": "default", "name": "Default"},
                         "id": "252a891c064648f38983f165e5aeb28d",
                         "name": "admin"},
             "catalog": [
                 {"endpoints": [
                     {"url": COMPUTE_URL.format(host=HOST, port=PORT, project_id=project_id),
                      "interface": "public",
                      "region": "RegionOne",
                      "region_id": "RegionOne",
                      "id": "41e9e3c05091494d83e471a9bf06f3ac"},
                     {"url": COMPUTE_URL.format(host=HOST, port=PORT, project_id=project_id),
                      "interface": "admin",
                      "region": "RegionOne",
                      "region_id": "RegionOne",
                      "id": "4ad8904c486c407b9ebbc379c58ea432"},
                     {"url": COMPUTE_URL.format(host=HOST, port=PORT, project_id=project_id),
                      "interface": "internal",
                      "region": "RegionOne",
                      "region_id": "RegionOne",
                      "id": "e39b209bbc0a4814a53be04fd761331c"}],
                  "type": "compute",
                  "id": "4a1bd1ae55854833870ad35fdf1f9be1",
                  "name": "nova"},
                 {"endpoints": [
                     {"url": NEUTRON_URL.format(host=HOST, port=PORT),
                      "interface": "internal",
                      "region": "RegionOne",
                      "region_id": "RegionOne",
                      "id": "9657ad081357459d9b6933df39d6aa90"},
                     {"url": NEUTRON_URL.format(host=HOST, port=PORT),
                      "interface": "admin",
                      "region": "RegionOne",
                      "region_id": "RegionOne",
                      "id": "c5a338861d2b4a609be30fdbf189b5c7"},
                     {"url": NEUTRON_URL.format(host=HOST, port=PORT),
                      "interface": "public",
                      "region": "RegionOne",
                      "region_id": "RegionOne",
                      "id": "dd3877984b2e4d49a951aa376c7580b2"}],
                  "type": "network",
                  "id": "d78d372c287a4681a0003819c0f97177",
                  "name": "neutron"},
             ],
             "user": {"domain": {"id": "default", "name": "Default"},
                      "id": "c95c5f5773864aacb5c09498a4e4ad0c",
                      "name": user},
             "audit_ids": ["DriuAdgyRoWcZG95-qpakw"],
             "issued_at": issued_at,
            }
    }

    return resp


@route('/nova/v2.1/<project_id>/images/detail', method='GET')
@route('/nova/v2.1/<project_id>/images', method='GET')
def nova_images_details(project_id):
    resp = {"images": [{"status": "ACTIVE",
                        "updated": "2016-12-05T22:30:29Z",
                        "id": "8591c262-032f-471e-b484-c23c7fbaff1d",
                        "OS-EXT-IMG-SIZE:size": 260899328,
                        "name": "trusty-server-cloudimg-amd64-disk1.img",
                        "created": "2016-12-05T22:29:35Z",
                        "minDisk": 20,
                        "progress": 100,
                        "minRam": 512,
                        "metadata": {"architecture": "amd64"}}]}
    return resp


@route('/nova/v2.1/<project_id>/images/<image_id>', method='GET')
def nova_image_get(project_id, image_id):
    resp = {"image":
            {"status": "ACTIVE",
             "updated": "2016-12-05T22:30:29Z",
             "links": [
                 {"href": "http://{host}:{port}/v2.1/{project_id}/images/{image_id}".format(host=HOST, port=PORT, project_id=project_id, image_id=image_id),
                  "rel": "self"},
                 {"href": "http://{host}:{port}/{project_id}/images/{image_id}".format(host=HOST, port=PORT, project_id=project_id, image_id=image_id),
                  "rel": "bookmark"},
                 {"href": "http://{host}:{port}/images/{image_id}".format(host=HOST, port=PORT, image_id=image_id),
                  "type": "application/vnd.openstack.image",
                  "rel": "alternate"}],
             "id": image_id,
             "OS-EXT-IMG-SIZE:size": 260899328,
             "name": "trusty-server-cloudimg-amd64-disk1.img",
             "created": "2016-12-05T22:29:35Z",
             "minDisk": 20,
             "progress": 100,
             "minRam": 512,
             "metadata": {"architecture": "amd64"}}}
    return resp


@route('/nova/v2.1/<project_id>/flavors/detail', method='GET')
def flavor_list_detail(project_id):
    resp = {"flavors":
            [{"name": "m1.small",
              "ram": 2048,
              "OS-FLV-DISABLED:disabled": False,
              "vcpus": 1,
              "swap": "",
              "os-flavor-access:is_public": True,
              "rxtx_factor": 1.0,
              "OS-FLV-EXT-DATA:ephemeral": 0,
              "disk": 20,
              "id": "2"}
            ]
    }
    return resp


@route('/nova/v2.1/<project_id>/flavors/<flavor_id:int>', method='GET')
def flavor_get(project_id, flavor_id):
    resp = {"flavor":
            {"name": "m1.small",
             "links": [
                 {"href": "http://{host}:{port}/v2.1/{project_id}/flavors/{flavor_id}".format(host=HOST, port=PORT, project_id=project_id, flavor_id=flavor_id),
                  "rel": "self"},
                 {"href": "http://{host}:{port}/{project_id}/flavors/{flavor_id}".format(host=HOST, port=PORT, project_id=project_id, flavor_id=flavor_id),
                  "rel": "bookmark"}],
             "ram": 2048,
             "OS-FLV-DISABLED:disabled": False,
             "vcpus": 1,
             "swap": "",
             "os-flavor-access:is_public": True,
             "rxtx_factor": 1.0,
             "OS-FLV-EXT-DATA:ephemeral": 0,
             "disk": 20,
             "id": int(flavor_id)}}
    return resp


@route('/neutron/v2.0/<network_id>', method='GET')
def network_get(network_id):
    return {'networks': 'yes please'}


@route('/nova/v2.1/<project_id>/servers', method='POST')
def server_create(project_id):
    response.status = 202
    return {"server": {"id": "fake-server-id"}}


@route('/')
def root():
    response.status = 300
    return {"versions":
            {"values":
             [
                 {"status": "stable",
                  "updated": "2016-04-04T00:00:00Z",
                  "media-types": [
                      {"base": "application/json",
                       "type": "application/vnd.openstack.identity-v3+json"}],
                  "id": "v3.6",
                  "links": [
                      {"href": "http://{host}:{port}/v3/".format(host=HOST, port=PORT),
                       "rel": "self"}]
                 },
                 {"status": "stable",
                  "updated": "2014-04-17T00:00:00Z",
                  "media-types": [
                      {"base": "application/json",
                       "type": "application/vnd.openstack.identity-v2.0+json"}],
                  "id": "v2.0",
                  "links": [
                      {"href": "http://{host}:{port}/v2.0/".format(host=HOST, port=PORT),
                       "rel": "self"},
                      {"href": "http://docs.openstack.org/",
                       "type": "text/html", "rel": "describedby"}]
                 }
             ]
            }
    }


run(host='0.0.0.0', port=PORT, debug=True)
