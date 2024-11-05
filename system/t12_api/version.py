from api_lib import APITest
import os
import json

class VersionAPITest(APITest):
    """
    GET /version
    """

    def check(self):
        actual = json.dumps(self.get("/api/version").json(), sort_keys=True)
        expected = json.dumps({'Commit': os.environ['APTLY_COMMIT'], 'Version': os.environ['APTLY_VERSION']}, sort_keys=True)
        self.check_equal(actual, expected)
