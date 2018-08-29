from site24x7beat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Site24x7beat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        site24x7beat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("site24x7beat is running"))
        exit_code = site24x7beat_proc.kill_and_wait()
        assert exit_code == 0
