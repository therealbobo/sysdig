/*
Copyright (C) 2013-2018 Draios Inc dba Sysdig.

This file is part of sysdig.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

#ifndef NOCURSESUI

class curses_table : 
	public curses_scrollable_list,
	public sinsp_chart
{
public:
	enum alignment
	{
		ALIGN_LEFT,
		ALIGN_RIGHT,
	};

	curses_table(sinsp_cursesui* parent, sinsp* inspector, chisel_table::tabletype type);
	~curses_table();

	void configure(chisel_table* table, 
		std::vector<int32_t>* colsizes, std::vector<std::string>* colnames);
	void update_data(std::vector<chisel_sample_row>* data, bool force_selection_change = false);
	void render(bool data_changed);
	chisel_table_action handle_input(int ch);
	void set_x_start(uint32_t x)
	{
		m_table_x_start = x;
	}
	void recreate_win(int h);
	void update_rowkey(int32_t row);
	void goto_row(int32_t row);
	bool get_position(int32_t* pos, int32_t* totlines, float* percent, bool* truncated);
	void follow_end();
	std::string get_field_val(std::string fldname);
	uint32_t get_data_size()
	{
		if(m_table != NULL)
		{
			return m_data->size();
		}
		else
		{
			return 0;
		}
	}

	chisel_table_field_storage m_last_key;
	bool m_drilled_up;
	bool m_selection_changed;
	MEVENT m_last_mevent;
	
private:
	alignment get_field_alignment(ppm_param_type type);
	void print_error(std::string wstr);
	void print_wait();
	void print_line_centered(std::string line, int32_t off = 0);

	sinsp* m_inspector;
	WINDOW* m_tblwin;
	sinsp_cursesui* m_parent;
	chisel_table* m_table;
	int32_t m_table_x_start;
	uint32_t m_table_y_start;
	uint32_t m_scrolloff_x;
	uint32_t m_colsizes[PT_MAX];
	std::vector<curses_table_column_info> m_legend;
	std::vector<chisel_sample_row>* m_data;
	sinsp_filter_check_reference* m_converter;
	std::vector<uint32_t> m_column_startx;
	char alignbuf[64];
	chisel_table::tabletype m_type;

	friend class curses_table_sidemenu;
	friend class sinsp_cursesui;        // for access to m_data in sinsp_cursesui::handle_input
};

#endif // NOCURSESUI
